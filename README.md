# arm64

The AArch64 instruction builder: registers, the ISA table, the encoder, and an in-memory object you can hand straight to an ELF, COFF, or Mach-O writer.

```go
m := arm64.NewModule()

t := m.Section(arm64.Text)
t.Label("main", arm64.Global, arm64.Func)
t.StpPre64(arm64.FP, arm64.LR, arm64.Mem64(arm64.SP).Pre(-16)) // a9bf7bfd — decided now
t.MovSp64(arm64.FP, arm64.SP)
t.Adrp(arm64.X0, arm64.Ref("msg"))
t.AddImm64(arm64.X0, arm64.X0, arm64.PageOff(arm64.Ref("msg")))
t.Bl(arm64.Ref("puts"))                                        // hole + reference, no PLT distinction to make
t.MovzImm32(arm64.W0, 0)
t.LdpPost64(arm64.FP, arm64.LR, arm64.Mem64(arm64.SP).Post(16))
t.Ret()
t.EndLabel("main")

m.Extern("puts")

o, err := m.Finalize()
if err != nil {
    log.Fatal(err)
}
elf.Write(f, o)
```

## Install

```sh
go get github.com/vertex-language/arm64
```

Container writers are leaf packages, so a build only pays for the format it emits:

```go
import "github.com/vertex-language/arm64/obj/elf"   // *obj.Object → ELF ET_REL
import "github.com/vertex-language/arm64/obj/macho" // *obj.Object → Mach-O MH_OBJECT
import "github.com/vertex-language/arm64/obj/pe"    // *obj.Object → COFF .obj
```

All three writers are in the tree and tested against tools this project did not write: `obj/elf`'s output is checked with `objdump`, `obj/pe`'s the same way, and `obj/macho`'s is linked with the host's own `clang`/`ld` and *executed* — on Apple Silicon that is a native link, no cross-toolchain involved. See [Writers](#writers).

## Targets

One architecture in, three containers out — and unlike its amd64 sibling, this is the architecture native to two of the three:

| | Container | Writer | Links in-tree? |
|---|---|---|---|
| **Linux, BSD, and other SysV** | ELF `ET_REL`, `ELFCLASS64`, `EM_AARCH64` | `arm64/obj/elf` | Not yet — `elf/link` resolves its machine through a registered backend and has none for AArch64 yet. `elf` itself now carries the `R_AARCH64_*` table (`reloc_aarch64.go`) this writer needs, added alongside this package. |
| **Apple** | Mach-O `MH_OBJECT`, `CPU_TYPE_ARM64` | `arm64/obj/macho` | Yes — `macho/arm64` is the backend `macho/link` already ships, because Apple Silicon *is* arm64 and that backend is not a cross-compilation add-on. |
| **Windows** | COFF `.obj`, `IMAGE_FILE_MACHINE_ARM64` | `arm64/obj/pe` | Not yet — `pe/link` has no ARM64 backend registered yet, though `pe/reloc_aarch64.go` carries the relocation vocabulary. |

Mach-O is the one container where this module's output is proven to run: [Writers](#writers) links a produced object with the host's linker and executes it as part of the test suite. ELF and PE are proven the way a disassembler proves an object — every section, symbol and relocation reads back correctly under `objdump` — but not linked in this tree, because doing that needs a cross-linker (`aarch64-linux-gnu-ld`, `lld`, or `link.exe`) this development environment does not have.

This module never learns what a linker is. Nothing here imports `link`, `backend`, or `image` from any of them.

## Contents

- [Targets](#targets)
- [Package map](#package-map)
- [Where `obj` lives](#where-obj-lives)
- [Quick start](#quick-start)
  - [Build a module](#build-a-module)
  - [Read the finished object](#read-the-finished-object)
  - [Lower from an IR](#lower-from-an-ir)
- [Instructions](#instructions)
- [Registers and widths](#registers-and-widths)
- [Sections and data](#sections-and-data)
- [Memory operands](#memory-operands)
- [Symbols](#symbols)
- [References](#references)
- [Features](#features)
- [Errors](#errors)
- [Writers](#writers)
- [How it's put together](#how-its-put-together)
- [Known limitations](#known-limitations)
- [License](#license)

## Package map

| Package | Purpose |
|---|---|
| `arm64` | The builder. `Module`, `Section`, typed helpers, `Emit`. This is what a frontend imports. |
| `arm64/reg` | `X0`..`X30`, `W0`..`W30`, `SP`, `V0`..`V31` and their scalar views, `Z0`..`Z31`, `P0`..`P15`, `Sys`. Imports nothing. |
| `arm64/operand` | `Imm`, `Mem`, `Label`, `SymRef`, `AddrRef`, `Cond`, `ShiftOp`, `ExtendOp`, `Barrier`, `PrfOp`. Imports `reg` and `obj`. |
| `arm64/feature` | The `Armv8-A`..`Armv9.5-A` levels, the orthogonal extensions, `ParseFeatures`. Stands alone. |
| `arm64/obj` | The vocabulary and the finished artifact: `Object`, `Section`, `Symbol`, `Reference`, `RefKind`, `SectionKind`, `Arch`, `Error`. Imports nothing. |
| `arm64/obj/elf` | `*obj.Object` → ELF relocatable object, via `github.com/vertex-language/elf`. |
| `arm64/obj/macho` | `*obj.Object` → Mach-O relocatable object, via `github.com/vertex-language/macho`. |
| `arm64/obj/pe` | `*obj.Object` → COFF relocatable object, via `github.com/vertex-language/pe`. |
| `arm64/internal/isa` | The form table: classes, opcodes, encodings, gates, `Resolve`, `ByHelper`. |
| `arm64/internal/encode` | form + operands → one word + fixups; the bit-field placement and the immediate arithmetic. |

Two directions, no cycles. `reg` imports nothing. `operand` sits on `reg` and on `arm64/obj`: it needs `reg` for the registers an address is built from, and `obj` for `RefKind`, because a reference's link semantics are a fact about the artifact and belong in the artifact's vocabulary. **`RefKind` is declared once, in `obj`, and aliased upward.** `arm64.RefAdrPage21`, `operand.RefAdrPage21` and `obj.RefAdrPage21` are the same constant, and no conversion exists anywhere in the tree.

`isa` and `encode` are `internal/` because they are implementation: the typed helpers and `Emit` are the instruction surface, and nothing a caller writes holds an `isa` or `encode` type. Encoder failures still reach you — through `*obj.Error` as a sentinel, a message, and notes. Data, not internal types.

### File layout

Each public package is one concept per file, and the instruction surface is split by tranche so a future `inst_sve.go` would sit opposite `isa/table_sve.go`'s row builder:

```
arm64/
  module.go section.go section_kind.go symbol.go error.go reloc.go
  alias.go build.go util.go emit.go inst.go
  inst_arith.go inst_alu.go inst_mov.go inst_bitfield.go inst_shift.go
  inst_cond.go inst_dp1.go inst_branch.go inst_loadstore.go inst_misc.go
  reg/       reg.go gpr.go vec.go sve.go sys.go save.go dwarf.go name.go
  operand/   operand.go imm.go mem.go sym.go shift.go extend.go cond.go
             barrier.go bitmask.go prfop.go
  feature/   feature.go level.go parse.go
  obj/       arch.go section_kind.go symbol.go error.go ref.go object.go
  internal/
    isa/     class.go form.go slot.go arg.go build.go alias.go isa.go
             resolve.go table_base.go
    encode/  encode.go operand.go imm.go fixup.go field.go nop.go error.go
```

Not landed yet, and named here because the shape is decided: `internal/isa/table_sve.go` and `table_sme.go`, each with an `inst_sve.go`/`inst_sme.go` opposite it — the `Z`, `P` and `Ffr` registers `reg/sve.go` already declares are waiting on that table, not on a register file that does not exist yet. And the three writers under `obj/`, each `write.go` translating this module's `RefKind` into its container's own relocation numbers.

## Where `obj` lives

`obj` is a package of this module, at `github.com/vertex-language/arm64/obj`. It imports nothing — not `reg`, not `operand`, nothing outside the standard library — which is what puts it at the bottom of the import graph beside `reg`, and what lets the three writers sit *under* it as `obj/elf`, `obj/macho` and `obj/pe` without a cycle. The amd64 and i386 modules have the same shape at their own `obj`, and the three are siblings rather than one shared dependency: there is no `github.com/vertex-language/obj` module and nothing here imports one.

That is a duplicated vocabulary, and it is duplicated on purpose. A shared module is a shared release cadence, and none of the three builders is finished enough to want one. What is not duplicated is the *shape* — `Arch`, `SectionKind`, `Symbol`, `Reference`, `RefKind`, `Error` appear in all three with the same fields doing the same work, so a diagnostic formatter or an object-inspection tool written against one reads the other two with no surprises.

`RefKind` is where this module's `obj` earns its keep rather than just copying amd64's. x86-64 materializes an address in one instruction and names one relocation; AArch64 materializes one in two — `ADRP` for the page, then an `ADD` or a load for the offset within it — and needs a kind for each half, plus a kind for the GOT-indirect version of each half, plus three more pairs for the three TLS models. `obj.RefKind` names all of it: `RefAdrPage21` and `RefAddAbsLo12` are the two halves of the position-independent pair, `RefAdrGotPage21` and `RefLd64GotLo12` are the GOT-indirect pair, and the branch fields — `RefCall26`, `RefJump26`, `RefCondBr19`, `RefTstBr14` — are named for the field each belongs to rather than folded into one generic "branch" kind, because a linker relaxing a call needs to know which one it is looking at. See [References](#references) for the full list and the reasoning behind each group.

## Quick start

### Build a module

```go
m := arm64.NewModule()

t := m.Section(arm64.Text)
t.Align(16)
t.Label("main", arm64.Global, arm64.Func)
t.StpPre64(arm64.FP, arm64.LR, arm64.Mem64(arm64.SP).Pre(-16))
t.MovSp64(arm64.FP, arm64.SP)
t.Adrp(arm64.X0, arm64.Ref("msg"))
t.AddImm64(arm64.X0, arm64.X0, arm64.PageOff(arm64.Ref("msg")))
t.Bl(arm64.Ref("puts"))
t.MovzImm32(arm64.W0, 0)
t.LdpPost64(arm64.FP, arm64.LR, arm64.Mem64(arm64.SP).Post(16))
t.Ret()
t.EndLabel("main")

r := m.Section(arm64.ROData)
r.Label("msg", arm64.Local, arm64.ObjectSym)
r.Asciz("hello\n")

m.Extern("puts")

o, err := m.Finalize()
if err != nil {
    log.Fatal(err)
}
```

`Finalize` folds every same-section fixup into its instruction word or its data hole, closes symbol sizes, resolves aliases and visibility, verifies that every surviving reference names something, and returns an `*obj.Object` — immutable, pure data, safe to write more than once, in more than one format, from more than one goroutine.

The builder is spent afterwards. Every call on `m` is a no-op recording `ErrFinalized`, and asking for a section that was never created is refused the same way: you get a usable handle whose every call is a no-op, and it is not added to the module, so `Sections()` is as frozen as the bytes are. `Finalize` is idempotent — call it again and you get the same object and the same error, not a second pass.

Section kinds and symbol attributes are `obj`'s constants, re-exported at the root. `arm64.Text` and `obj.Text` are the same value, so a lowering that never imports `obj` still spells them and a downstream consumer never converts. Every example here uses the root spelling.

### Read the finished object

```go
for _, s := range o.Sections() {
    s.Index()     // its position, which is what a symbol's Section names
    s.Kind()      // arm64.Text, arm64.Data, arm64.ROData, arm64.BSS
    s.Name()      // ".text", or whatever SectionNamed was given
    s.Align()     // the largest alignment the builder asked for, at least 1
    s.Size()      // length in bytes
    s.Bytes()     // finished machine code, same-section labels patched
    s.Refs()      // []obj.Reference — the holes a linker fills
    s.Symbols()   // filtered view over the object's one symbol table
}

for _, sym := range o.Symbols() {
    sym.Name, sym.Section, sym.Offset, sym.Size
    sym.Binding, sym.Type, sym.Visibility
    sym.Defined()  // Section is -1 when undefined
}

o.Symbol("main")        // by name; identity is the name, module-wide
o.SectionNamed(".text")
o.SectionAt(0)
```

`Bytes()` returns a copy, and so does every other slice accessor. The symbol table is module-level and in definition order, which is what every object writer wants.

### Lower from an IR

The builder is the only AArch64-specific type a frontend touches. Write your instruction selection against `*arm64.Section`, and everything downstream — object emission, linking, testing — is `obj.Object`:

```go
func (l *Lowering) block(t *arm64.Section, b *ir.Block) {
    t.Label(b.Name)                       // bare: a branch target, not a symbol
    for _, in := range b.Insts {
        switch in.Op {
        case ir.Add:
            t.AddShifted64(l.reg(in.Dst), l.reg(in.A), l.reg(in.B))
        case ir.Load:
            t.LdrImm64(l.reg(in.Dst), l.mem(in.Src))
        case ir.Call:
            t.Bl(arm64.Ref(in.Sym))
        case ir.Br:
            t.B(arm64.Label(in.Target))
        }
    }
}
```

No error checking in that loop, and that is the point: errors are sticky and first-wins, so the first failure positions itself and every later call is a no-op. Check once at `Finalize`, or call `m.Err()` to bail out of a long build early.

Calling conventions are not in this package. AAPCS64 puts the first eight integer arguments in X0 through X7, returns small aggregates the same way, and gives you no red zone; Apple's variant keeps the general shape but tightens the stack-argument alignment and drops the frame-pointer-chain requirement on leaf functions. Which one you are lowering for is a fact about your target, and this module does not have one — it has an instruction set. `Ret` emits `d65f03c0` and asks no questions about who set up the frame.

## Instructions

### Typed helpers — the primary surface

One method per declared form, named `MnemonicClassClass`:

```go
t.AddImm64(arm64.SP, arm64.SP, 16)                    // 64-bit add, immediate
t.SubShifted64(arm64.X0, arm64.X1, arm64.X2, arm64.Shifted(arm64.LSL, 3))
t.LdrImm64(arm64.X0, arm64.Mem64(arm64.SP).Off(16))
t.Csel64(arm64.X0, arm64.X1, arm64.X2, arm64.GT)
t.Cbz64(arm64.X0, arm64.Label("done"))
```

The operand classes are the parameter types, so a width or class mismatch is a compile error: `AddShifted64(arm64.W0, …)` does not build, because `W0` is a `reg.W` and the form wants a `reg.X`. An isel bug is a red squiggle rather than a runtime `ErrForm`.

A helper pins its form — exactly the encoding you named — and the diagnostics follow from that. A helper checks operand *kinds* only and hands values to the encoder: the wrong kind of operand is `ErrForm`; a value that does not fit the field the form pins is `ErrRange`, with the field width and range in the error's notes and the encoder's error as the cause. A logical-immediate constant with no rotated-run-of-ones encoding is `ErrBitmask` rather than `ErrRange`, because the fix is different in kind — the value has to be materialized into a register, not made smaller. Gated helpers still gate: a gated helper on a module without the level or extension fails at the call with `ErrFeature`, naming the gate.

`inst_*.go` binds each helper to its form **by name lookup, not by table index**. Appending rows breaks nothing, and a removed or renamed form panics at program start naming the missing form, rather than silently binding to the wrong row or failing halfway through someone's code generation. Two helpers binding the same name is a duplicate Go identifier and so a compile error, which is the earliest any of these can fail.

Naming conventions beyond the class spelling:

- **Fixed operands are in the name, not the parameters.** `Cset64(rd, cond)` is CSINC with both sources pinned to the zero register — there is no field left to put another register in, so there is no parameter for one.
- **Branch and address targets take one operand type, `TargetOp`.** A bare `Label("loop")` folds at `Finalize` with no relocation; a `SymRef` from `Ref("f")`, optionally wrapped in `Page`/`GotPage`/`PageOff`/`GotPageOff`, survives into `Refs()`. Which one you hand `Adrp` or `Bl` decides same-section-fold versus relocation — there is no second method name to choose between, because the operand already says which you meant.
- **Writeback is a separate form, not a flag.** `Stp64` takes a plain-offset `Mem`; `StpPre64` and `LdpPost64` take a pre- or post-indexed one and refuse a `Mem` built the other way, naming which addressing mode the form actually has. There is no single `Stp64(..., writeback bool)` — a boolean parameter would still leave the operand and the form disagreeing about the address until the call actually runs.
- **The bitfield family takes `immr`/`imms` or a single `shift`, matching the table exactly.** `Ubfm64`, `Sbfm64` and `Bfm64` take both fields, because that is the general encoding; `LslImm64`, `LsrImm64` and `AsrImm64` take one shift amount, because they are aliases that compute the other field for you.
- **Both documented spellings exist and both emit the same bytes.** `Cmp` is `Subs` with the destination pinned to the zero register; `Mov` is `Orr` or `Add` depending on which alias matched. They are separate forms with an alias relation checked at table build time, so a listing says which name the caller used and the bytes say what the silicon does.

### Emit — the escape hatch

```go
t.Emit("add", arm64.X0, arm64.X1, arm64.X2)
t.Emit("b", arm64.Label("done"))
```

Runtime form resolution for table-driven emission, where the mnemonic is data. If you know the instruction at compile time, use the typed helper.

There is no shortest-form search here and no tie to break, unlike the x86 modules in this tree. Every AArch64 instruction is one word, so two forms of a mnemonic that could both accept one operand list is a table-build-time error — `internal/isa`'s `checkTable` panics naming the collision — rather than a choice `Resolve` has to make at run time.

`Emit` does not yet discriminate a memory operand's addressing form — plain offset, pre-indexed, post-indexed — the way the typed helpers' pinned forms do; see [Known limitations](#known-limitations). For `Stp`/`Ldp`'s writeback forms and anything with a `Mem` operand, prefer the typed helper.

A mnemonic the table does not have is `ErrForm` saying so by name. Operands no form accepts is `ErrForm` naming the operands. Operands that matched a form the feature set does not permit is `ErrFeature` naming the gates that would allow it.

## Registers and widths

```go
arm64.X0   arm64.W0    // 64-bit and 32-bit views of the same register
arm64.SP   arm64.WSP   // register 31, read as the stack pointer
arm64.XZR  arm64.WZR   // register 31, read as the zero register
arm64.V0   arm64.Q0    arm64.D0  arm64.S0  arm64.H0  arm64.B0
```

Thirty-one general-purpose registers in two widths each, thirty-two vector registers in six widths, thirty-two scalable vector registers and sixteen predicates declared and waiting on the SVE tranche. Each width is a distinct Go type, which is what makes `AddShifted64(arm64.W0, …)` fail to compile.

One fact about this architecture is load-bearing enough to leak through the surface on purpose: **register 31 is two different registers, and which one a slot reads depends on the form, not on anything the number itself carries.** `AddImm64`'s destination and first source read 31 as `SP`; `AddShifted64`'s registers read it as `XZR`. This package spells the two as different Go types — `reg.Xsp` and `reg.X` — precisely so that handing `XZR` to a slot that means `SP`, or `SP` to a slot that means `XZR`, is a checked mismatch rather than a silent one. Where a slot accepts either reading of a numbered register (add's destination, for instance, which is `Xsp` but also takes a plain `X` that is not `XZR`), the parameter is a documented `any` — `RegSP64`, `RegSP32` — the same pattern the x86 modules in this tree use for their own union-shaped operand slots.

There is no REX-equivalent and nothing to leak by omission the way `AH` does on amd64: every general-purpose register is reachable in every form that names its class, at every width the form declares, with no encoding that reaches one register only when another is absent.

## Sections and data

There is one section per kind under the conventional name, and `Section(kind)` creates it on first use:

```go
t := m.Section(arm64.Text)        // ".text"
```

For anything outside those four names — debug info, unwind tables, a custom note section — `SectionNamed` takes the name and the load-time kind it behaves as:

```go
d := m.SectionNamed(".debug_line", arm64.Data)
d.Data(dwarfBytes)
```

`Section(k)` is exactly `SectionNamed(k.String(), k)`, so the two cannot disagree about `.text`. A name asked for twice returns the same section; a name asked for with a different kind than it was created with is `ErrDuplicate`, because a section that is code on Tuesday and data on Wednesday is a bug with a delayed fuse.

```go
r.Byte(0x90)
r.Long(0xdeadbeef)         // little-endian, 4 bytes
r.Quad(1)                  // little-endian, 8 bytes
r.Ascii("no terminator")
r.Asciz("terminated")
r.Zero(64)
r.Data(blob)               // raw bytes
r.Ref("f", arm64.RefAbs64) // 8-byte hole + a relocation
r.LabelRef("case_3")       // 8-byte hole, patched at Finalize
```

`Ref` and `LabelRef` are the data-side twins of what an instruction operand does. They exist because this package refuses to build your vtables and jump tables and must not make them unbuildable.

`Align(n)` pads a code section with `D503201F` — the architecture's one canonical no-op, four bytes, always — and a data section with zeros. `n` must be a power of two, and in a code section a multiple of four: every instruction here is one word, and an alignment that would strand a partial one is `ErrAlign` rather than rounded. The largest `n` a section sees also becomes `Section.Align()`, which is what the object writer stamps on the section header.

`Offset()` is the current end of the section: the offset the next word will land at, the value a `Label` placed now would name. It is exported because those tables are yours to build, and building them requires knowing where you are.

## Memory operands

One constructor per access width, refined by chain methods that keep the width's type:

```go
arm64.Mem64(arm64.X1)                                   // based, no offset:   [x1]
arm64.Mem64(arm64.X1).Off(8)                             // scaled offset:      [x1, #8]
arm64.Mem64(arm64.SP).Pre(-16)                           // pre-indexed:        [sp, #-16]!
arm64.Mem64(arm64.SP).Post(16)                           // post-indexed:       [sp], #16
arm64.Mem64(arm64.X1).Indexed(arm64.X2, arm64.ExtLSL, 3) // register offset:    [x1, x2, lsl #3]
arm64.Mem64(arm64.X1).Off(arm64.PageOff(arm64.Ref("v"))) // symbolic page offset, the ADRP/LDR pair's second half
```

The scaled/unscaled distinction `internal/isa` carries is deliberately absent from `Mem` itself: `[x1, #8]` is the same operand whether it lands in `LdrImm64` or `LdurImm64`, because those are two mnemonics you named and picking between them on the strength of whether the offset happens to divide evenly is exactly the instruction selection this tree does not do.

A displacement that is a symbol reference must state which half of the address it means — `PageOff` or `GotPageOff` — because the low-twelve half of an address is a role you have to write down, and guessing would be inventing the one thing a reader most needs to see. A bare `Label` or `SymRef` handed to `.Off` is refused for that reason.

`Indexed`'s index register is a `reg.X` or a `reg.W` and its extend states which: `UXTW`/`SXTW` for a 32-bit index, `LSL`/`SXTX` for a 64-bit one. Getting the two mismatched — a `W` index with `SXTX`, an `X` index with `UXTW` — is `ErrOperand` naming both.

Construction errors — an index register on an address with no room for one, a shift amount that is not zero or the log of the access width, a symbolic offset on a writeback address, which has nowhere for a relocation to go — are sticky on the operand and surface at the instruction that uses it, positioned, under `ErrOperand`. A builder chain is not followed by a run of error checks, and the diagnostic still points at the instruction rather than at the encoder.

## Symbols

```go
t.Label("loop")                              // bare: a branch target, this section only
t.Label("main", arm64.Global, arm64.Func)    // attributed: a symbol, size range opens
t.EndLabel("main")                           // closes it
m.Extern("puts")                             // undefined, Section == -1
m.Alias("_main", "main")                     // second name, same offset
m.SetVisibility("main", arm64.Hidden)
```

**A bare `Label` is not a symbol.** It names an offset in this section's namespace, gets folded into whatever field names it at `Finalize`, and leaves no trace in the symbol table. Any attribute — a `Binding`, a `SymbolType`, a `Visibility` — promotes it into `Symbols()`. That promotion is not just bookkeeping here: a page or GOT reference to a label depends on where the section finally loads, which nothing at this layer assigns, so `Adrp`/`Page`/`GotPage` against an unpromoted label is `ErrUndefined` naming the label and telling you to promote it — only a bare branch or `Adr`'s direct target can fold same-section.

`Size` closes at `EndLabel` if you call it, at the next symbol in the same section if you don't, and at the section end otherwise. It is stated rather than guessed because a zero-size function symbol defeats dead-stripping in every linker that does it, and the next-symbol fallback is a guess, so prefer `EndLabel` for anything you care about.

`Alias` and `SetVisibility` resolve at `Finalize`, so the order of the call and the `Label` it names does not matter.

Symbol identity is the name, module-wide. A duplicate definition is `ErrDuplicate` at the second `Label`, naming the first one's section and offset. Nothing here adds a leading underscore; if you are targeting Mach-O and want `_main`, spell it.

## References

```go
arm64.Ref("puts")                                  // a plain reference; the caller does not insist on a kind
arm64.Ref("f", arm64.RefCall26)                     // insist on a call-family relocation
arm64.PageOff(arm64.Ref("msg"))                     // the :lo12: half of an ADRP/ADD pair
arm64.GotPage(arm64.Ref("errno"))                   // the page of a GOT-indirect load
```

```
RefAbs64  RefAbs32  RefAbs16                          absolute, that many bytes — data only
RefPrel64  RefPrel32  RefPrel16                        PC-relative — data only
RefCall26  RefJump26  RefCondBr19  RefTstBr14          the four branch field widths
RefAdrPrel21  RefAdrPage21                             ADR's direct field, ADRP's page field
RefAddAbsLo12
RefLdSt8AbsLo12  RefLdSt16AbsLo12  RefLdSt32AbsLo12
RefLdSt64AbsLo12  RefLdSt128AbsLo12                     the page-offset half, by access width
RefAdrGotPage21  RefLd64GotLo12                        the GOT-indirect pair
RefTlsGdAdrPage21  RefTlsGdAddLo12                      general-dynamic TLS
RefTlsIeAdrGottprelPage21  RefTlsIeLd64GottprelLo12     initial-exec TLS
RefTlsLeAddTprelHi12  RefTlsLeAddTprelLo12              local-exec TLS
RefTLV                                                  Mach-O thread-local variable
RefSize32  RefSize64  RefSecRel32  RefSecIdx            COFF- and size-relative forms
```

Names follow the ELF psABI's own relocation names — `R_AARCH64_ADR_PREL_PG_HI21` becomes `RefAdrPage21` — because that vocabulary is what every AArch64 ABI document, disassembler, and linker error message already uses, and a second naming scheme for the same handful of ideas would only cost a reader the translation.

That list is the *union* of what the three containers can express, not the intersection, for the same reason the x86 modules in this tree make the same choice: an intersection would drop the page/page-offset pairing that is how every position-independent AArch64 address is built, and a per-container kind set would mean a lowering picks its container before it picks its instructions. Every writer states what it cannot do, and `ErrRefKind` names the kind and the offset when a kind meets a container with no answer for it.

The `LdSt*AbsLo12` family is five kinds rather than one because the immediate is scaled by the access width — `LDR` (64-bit) shifts the page offset right by three before checking it fits, `LDRB` does not shift at all — and a linker applying the wrong one computes the wrong address silently rather than refusing. The kind is chosen from the load or store's own access width, not stated by the caller: `LdrImm64(rt, Mem64(base).Off(PageOff(sym)))` picks `RefLdSt64AbsLo12` because the `Mem64` already said the width.

`RefCall26` and `RefJump26` share one field — a 26-bit word-scaled displacement — and differ only in what a linker may do at the far end: `RefJump26` (a plain `B`) may be redirected to a long-branch thunk as a tail call, where `RefCall26` (`BL`) must return through the same instruction that reached it. `Bl` and `B` each bind to their own kind for exactly that reason, and naming the wrong one on an explicit `Ref` would tell a linker the wrong thing about a return address.

### The one identity this format needs and x86 does not

```go
type Reference struct {
    Offset int      // where the hole starts, section-relative
    Size   int       // 2, 4 or 8 for a data kind; 4 — one instruction word — for every other kind
    PCRel  bool
    Adjust int64    // always 0 on this architecture
    Sym    string
    Kind   RefKind
    Addend int64    // logical addend, never adjusted for the field
}
```

`Adjust` exists on the other two architectures in this tree for a PC-relative field that does not resolve against its own start — x86-64's displacement resolves against the end of the instruction, wherever that falls, because the field is not always the last thing in it. Nothing on AArch64 works that way: every PC-relative field here, `ADRP`'s included, resolves against the address of its own instruction word. A `Reference` this package builds always carries `Adjust == 0`. The field is kept rather than dropped so a writer shared textually with the other `obj` packages does not need a special case for the one architecture that never sets it.

`Size` is the field itself for the data kinds — a `RefAbs64` pointer is eight bytes wherever it appears — and always four, one instruction word, for every kind that names a bit-field inside an instruction. The bit position and width within that word are a fixed fact about the kind, which a writer looks up rather than a number `Reference` could usefully carry; that is what makes the same `Reference` shape describe both a data-section pointer and a branch's 26-bit field without a variant for either.

## Features

```go
m := arm64.NewModule(arm64.WithFeatures(
    arm64.Armv8_2A.Plus(arm64.PAuth, arm64.BTI),
))
```

Default is Armv8-A, the baseline, with no extensions. The feature set is fixed at construction and nothing about it is configurable after: a gate that changed mid-module would make already-emitted diagnostics unfalsifiable.

Two different things live in the vocabulary and are deliberately not flattened together. A **level** is a point on the cumulative Armv8.x-A / Armv9.x-A ladder Arm's own architecture reference defines — Armv8.1-A adds CRC, LSE and RDMA over the baseline; Armv9-A is Armv8.5-A plus SVE and SVE2 — and a level is a bundle rather than an instruction group. A **feature** is an orthogonal extension with an ID-register field of its own and no level that requires it: AES, SHA2, SHA3, SM4, PAuth, BTI, MemTag, and the SVE and SME sub-extensions past the level that first introduces the family.

Sets are closed under requirements in both directions: adding SVE2 brings SVE, removing SVE drops SVE2 and every SVE2 sub-extension, because a set holding SVE2 but not SVE describes no silicon.

`ParseFeatures` resolves the spellings the world writes, against a starting set:

```go
arm64.ParseFeatures("armv8.6-a+sve+sve2")       // exact: level plus extensions
arm64.ParseFeatures("+bti,-pauth")              // adjust: applied left to right, from Baseline
```

`Set.String()` prints the canonical spelling and `ParseFeatures` accepts it back, so the two round-trip. Removing a feature a level requires demotes the level and prints as the lower level plus what survives; there is no way to hold a set that says `armv9-a` and not `sve`.

`feature.Levels`, `Level.Set`, `Level.Plus`, `Level.Minus` and `Decompose` are there so a driver can print its own `--help`. The data is the library's; the formatting is the consumer's.

## Errors

```
ErrFeature  ErrForm  ErrOperand  ErrDuplicate  ErrUndefined
ErrRange  ErrBitmask  ErrAlign  ErrFinalized  ErrRefKind  ErrSectionName
```

Each names one failure and only one:

- `ErrForm` — the operands are the wrong *kinds* for the form, or no form of that name exists.
- `ErrOperand` — the operands are the right kinds and one of them was built wrong. `SP` where a slot means the zero register, a shift amount out of range for the width, an index register whose extend reads the wrong width. It is its own sentinel because "no matching form" sends a caller hunting through the ISA table for a row that exists.
- `ErrRange` — a value does not fit its pinned field. Both the immediate at the call site and the branch displacement at `Finalize`, both carrying the field width and reachable range in the notes. There is no branch relaxation and no silent form substitution: the failure is loud instead of the bytes being different.
- `ErrBitmask` — a constant handed to a logical-immediate field — `AndImm64`, `OrrImm64`, `EorImm64` — is not representable as one: not every value near the field's width is a rotated run of ones replicated to fill the register. The fix is different from `ErrRange`'s: materialize the constant into a register instead of trying to narrow it.
- `ErrRefKind` — the writer has no relocation for a kind, or a literal load (`LdrLit64`) was given a symbol rather than a same-section label, which has no encoding in any of the three container formats. Only ever comes from a writer, or from the one check this package makes on its own behalf.
- `ErrSectionName` — the writer cannot place a custom section. Only ever comes from `arm64/obj/macho`.

The concrete type is `*obj.Error`, with `Arch`, `Section`, `Offset`, `Context`, `Sentinel`, `Cause` and `Notes`. One error type across the builder and all three writers, and across every architecture in the tree, so tooling that formats a diagnostic is written once:

```
arm64 .text+0x14: displacement 1048580 does not fit 19 signed bits: value out of range
```

Errors are sticky and first-wins: every builder call after a failure is a no-op, and `Finalize` surfaces the first one, positioned. `Module.Err()` returns the same error `Finalize` will.

`errors.Is` works against every sentinel. Where a resolver or encoder error is the cause, it joins the chain — `Unwrap` returns both the sentinel and the cause — but its type is internal; anything a caller might need from it is restated in `Notes` as text.

## Writers

All three are in the tree. Each is one exported `Write(w io.Writer, o *obj.Object, opts ...Options) error` and an `Options` struct for what the assembler has no opinion about. No `Target` field on any of them: the object names `obj.ArchARM64` and everything each container's header needs follows from it.

```go
elf.Write(f, o)
elf.Write(f, o, elf.Options{OSABI: elfcore.ELFOSABI_GNU, Comment: "vertex 0.1"})

macho.Write(f, o, macho.Options{Platform: machocore.PlatformMacOS, MinOS: "11.0"})

pe.Write(f, o)
pe.Write(f, o, pe.Options{File: "hello.c", Directives: []pe.Directive{{Name: "DEFAULTLIB", Value: "msvcrt"}}})
```

Each is tested against a tool this tree did not write, not against its own idea of what it produced:

- **`arm64/obj/elf`** is RELA throughout, the same model as amd64's ELF writer: the addend lives in the relocation entry and the section bytes go through untouched. `elf_test.go` writes an object and checks its sections, symbols and `R_AARCH64_*` relocations under `objdump -r`, `-t` and `-d`.
- **`arm64/obj/macho`** is the one writer this test suite proves end to end: `macho_test.go` writes an object, links it with the host's own `clang`, and *runs the result*, checking the program's actual output. On an Apple Silicon host — the only one the test runs on — that is a native link with no cross-toolchain and nothing to fake. Most of AArch64's Mach-O relocations carry no addend at all: the ADRP/ADD/LDR pair's fields are the page and the page offset, computed by the linker from the symbol alone, with nothing for this writer to fold in. A nonzero `Addend` on one of those is `ErrRefKind` today — see [Known limitations](#known-limitations) — because expressing it needs a preceding `ARM64_RELOC_ADDEND` entry this writer does not emit yet. `RefAbs64`/`RefAbs32`/`RefAbs16` are the exception: `ARM64_RELOC_UNSIGNED` folds an addend into the field exactly the way x86-64's Mach-O writer does.
- **`arm64/obj/pe`** has the simplest relocation model of the three: ARM64 COFF relocations never pair (`RelocARM64.IsPair` and `TakesPair` both always report false), so every `RefKind` is one relocation record and nothing more. The `PAGEOFFSET_12L` type folds all five of `RefLdSt8AbsLo12` through `RefLdSt128AbsLo12` into one number — the psABI's own note is that the access-width scaling is a linker's job, decoded from the instruction, not a table's — which is simpler than ELF's and Mach-O's per-width families. `pe_test.go` checks the same three things the ELF test does, under `objdump`. There is no link-and-run counterpart: this development host has no ARM64 Windows linker.

`ROData` becomes `.rdata` for COFF, matching link.exe's default merge rules, the same rename amd64's PE writer makes; every other section keeps its ELF-spelled name across all three containers except Mach-O, which places by segment and section (`(__TEXT,__text)`, `(__TEXT,__const)`, and so on) rather than by name at all.

A `RefKind` a container has no relocation for is `ErrRefKind` from that writer, naming the kind, the symbol and the section offset — never a construction-time refusal, because the same object is legal for a different container. The GOT kinds and every TLS model are `ErrRefKind` from `obj/pe` (no GOT, no ELF-shaped TLS on this machine); every ELF TLS model except the plain descriptor case is `ErrRefKind` from `obj/macho` (`RefTLV` is the one thread-local kind that container answers for, and nothing above the reference layer builds the descriptor sequence yet regardless).

## How it's put together

- **The builder is concrete, the artifact is inert.** Typed helpers are methods on `*arm64.Section` because that is what makes a width mismatch a compile error; an interface or a generic builder would erase exactly the checking the surface exists for. Everything past `Finalize` is `obj.Object`: data with no methods that do anything but read.
- **`obj` imports nothing and knows no ISA.** It is the vocabulary the rest of the tree spells its types in, and the finished artifact the builder hands to a writer. That it imports nothing is what lets the writers live under it — `obj/elf`, `obj/macho`, `obj/pe` — while `operand` and `internal/encode` import it from above, with no cycle either way.
- **One declaration per concept, aliased upward.** `RefKind` is `obj`'s, aliased by `operand` and the root. `SectionKind` and the symbol attributes are `obj`'s, re-exported at the root. A value crossing a package line is never converted, only renamed.
- **Every instruction is one word, and the table says so up front.** There is no shortest-form search and no size estimator anywhere in this tree's AArch64 half, because there is nothing to estimate: `internal/encode.EncodeForm` always produces exactly four bytes plus zero or more fixups. `internal/isa.checkTable` panics at program start if two forms of one mnemonic could ever accept the same operand list — an ambiguity x86's variable-length encoding has to live with and this architecture's fixed width does not.
- **A relocation kind names a field, not a byte range.** x86's `Reference.Size` is how many bytes a writer touches; most of this architecture's kinds instead say which instruction the field belongs to, because the field is a run of bits inside one word rather than a byte-aligned range. `RefKind.Size` still answers 4 for those — one word — but the bit layout is a fact the kind's name carries, not a number in the struct.
- **The alias relation is checked, not assumed.** `CMP`, `CMN`, `TST`, `MOV` and the shift mnemonics are declared as narrower forms of `SUBS`, `ADDS`, `ANDS`, `ORR`/`ADD` and `UBFM`/`SBFM`, each with the fields it pins stated explicitly. `internal/isa.checkAliases` verifies at program start that a pinned field really is fixed in the alias's own word and mask — an alias whose pinned field were not in its mask would decode as the underlying instruction too, which is exactly the ambiguity the check exists to catch before any caller can hit it.
- **Two table-integrity checks run before any caller can be wrong.** `isa.checkTable` panics on two forms sharing a helper name or an identical operand signature; `isa.checkAliases` panics on an alias whose pinned fields do not hold. Both are questions about this tree's own data and should fail whatever the caller is, at program start, rather than at some caller's build.

## Known limitations

- **The table is the base integer, branch, and load/store instruction set plus scalar floating point and the atomics: 244 forms over 145 mnemonics.** `table_base.go` declares the ALU block (immediate, shifted-register, extended-register), the logical and bitfield families, move-wide and `ADR`/`ADRP`, the shift/multiply/divide group, conditional select, the unary data-processing group, the full branch family, scalar load/store including pairs and literals, the barrier and exception instructions, and `MRS`/`MSR` for the common system registers. `table_atomic.go` declares the acquire-release accesses, the load/store-exclusive pairs, and the acquire-release half of the LSE family. `table_float.go` declares scalar single and double precision: the arithmetic, the three-source fused multiply-adds, both min/max pairs, the four rounding modes, the width conversions, the conversions to and from both integer widths, `FMOV` between the register files, `FCMP`, `FCSEL`, and the FP loads and stores. `reg/sve.go` declares the `Z` and `P` types and `reg/vec.go` the vector-with-arrangement ones, and nothing in the table reaches those yet — SVE, SME, and Advanced SIMD are each a tranche of their own, the same way AVX and AVX-512 are named-but-unbuilt on the amd64 side of this tree.
- **`Emit`'s dynamic dispatch does not discriminate a memory operand's addressing form.** `Resolve` matches on operand class and access width, not on whether a `Mem` was built with `.Off`, `.Pre`, or `.Post` — so a runtime `Emit("stp", ...)` against a pre-indexed `Mem` is not guaranteed to reach `StpPre64` over `Stp64` the way a typed-helper call is by construction. The typed helpers are unaffected: each pins its own form directly and `internal/encode` refuses a `Mem` built the wrong way for it, by name, at the call. Prefer the typed helper for any writeback form until `Resolve` carries the same discrimination the encoder already does.
- **No unwind-table generation.** `.eh_frame` on SysV, `__unwind_info` and `__eh_frame` on Apple, `.pdata`/`.xdata` on Windows ARM64. All three are bytes: build the sections yourself with `SectionNamed` and `Data`.
- **No pointer authentication or branch-target-identification instruction support.** `PAuth` and `BTI` exist as features a module can gate on, but `PACIASP`, `AUTIASP`, `BTI c`, and the rest of that family are not in the table — a module built with those features enabled gates nothing yet, because nothing gated needs them.
- **No SVE, SME, or Advanced SIMD.** See above. The `S`, `D` and `Q` scalar views have instructions now; `H` and `B` do not, half precision being a separate `ftype` and a separate feature, and every SIMD-with-arrangement form remains a register type with nothing to reach it.
- **The atomics are the sequentially consistent ones.** `LDAR`/`STLR`, the four exclusive pairs, and the LSE family's acquire-release variants are declared; the relaxed, acquire-only and release-only spellings of the LSE forms (`LDADD`, `LDADDA`, `LDADDL` beside `LDADDAL`, and so on for each) are not. Every one is a two-bit change to a declared row, and none is there because nothing selects it yet — a row nothing selects is a row nothing tests. The exclusive pairs make every atomic expressible at any ordering regardless, LSE being the one-instruction form of the loop rather than the only way to write it.
- **No cross-section symbol differences and no COMDAT-equivalent section grouping.** Same gaps as the other two architectures in this tree, for the same reasons: a same-section `LabelDiff` (not yet ported here — see below) needs no relocation and covers the common case; the cross-section case is a shared-`obj`-vocabulary change that should be made once, deliberately, for all three architectures together.
- **`LabelDiff` and `Ascii`/`Asciz` byte-level table-building helpers exist on the other two architectures in this tree and are not yet on this one's `Section`.** The underlying mechanism — a same-section fixup with no relocation — is already how `foldLabel` resolves a branch; extending it to a plain difference of two labels is a small, well-understood addition rather than a design question.
- **No disassembler.** This module writes bytes. Reading them back is `elf/obj`, `pe/coff`, and `macho/obj`, which is where the writer tests get the other half of their round trip.
- **A nonzero `Reference.Addend` on an ADRP-family kind is `ErrRefKind` from `obj/macho`.** Expressing one needs a preceding `ARM64_RELOC_ADDEND` entry — `macho/obj.Writer.RelocPair` supports emitting the pair structurally, but this writer does not build one yet. `RefAbs64`/`RefAbs32`/`RefAbs16` are unaffected: `ARM64_RELOC_UNSIGNED` takes an implicit addend the same way it does on amd64.
- **`obj/elf` and `obj/pe` are proven by disassembly, not by linking, in this tree.** Both round-trip correctly under `objdump`, but neither is linked here: `elf/link` and `pe/link` have no registered ARM64 backend yet, and this development environment has no standalone `aarch64-linux-gnu-ld`, `lld`, or `link.exe` to hand an object to directly. `obj/macho` is the one writer proven by an actual link-and-run, because Apple Silicon supplies a native linker for it.
- **`AddExt32`/`AddExt64`/`SubsExt64`'s second source is one width, not either.** The extended-register add/sub forms accept a 32-bit or a 64-bit `Rm` depending on the extend applied to it — `UXTW`/`SXTW` read a `W`, `UXTX`/`SXTX`/`LSL` read an `X` — but the table declares a single register class per slot, so today `AddExt64`'s third operand is `ClassX` only and a `reg.W` is refused even under `UXTW`. Accepting both correctly needs the slot's accepted class to depend on a *later* operand, which the class system does not do yet; encoding a 32-bit-extended add into a 64-bit destination has no typed helper until it does.
- **Register-offset addressing has no encoding.** `[Xn, Xm]`, `[Xn, Xm, LSL #s]`, and `[Xn, Wm, SXTW #s]` are refused by `Mem.Validate`'s caller with `UnsupportedError`, naming the address and saying so — the table declares no `Rm`/option field for that addressing form on any load or store row yet. Every load/store helper reaches only the immediate-offset and literal forms.
- **Differential testing exists and covers most of the table, not all of it.** `TestDifferentialAgainstClang` (in `difftest_test.go`) assembles 243 of the table's cases with clang's integrated AArch64 assembler and checks the typed-helper output byte-for-byte against it, and it is what caught — and this package fixed — several real encoding bugs the self-consistency checks alone could not: `ImmKind` (the arithmetic behind logical immediates, move-wide, scaled offsets, and more) was declared throughout `internal/isa/slot.go` but had no builder method ever wired it to a row, so every one of those forms silently fell back to a plain range-checked placement; a bit-trick in `operand.EncodeBitmask` mis-detected which constants are valid logical immediates; `MRS`/`MSR` had no dispatch case in `internal/encode` at all; a shifted-register form's shift *amount* was validated but never placed in the word; `LSL`/`LSR`/`ASR` (immediate) never computed the `UBFM`/`SBFM` fields they alias; and `Cset64` encoded the caller's condition uninverted. The suite is skipped, not failed, when clang is not on `PATH`. Every scalar floating-point and atomic form is covered, which is how they were written: clang was asked for the encoding of each one and the table row derived from its answer, so the suite is the derivation checked rather than a check added afterwards. It assembles at `.arch armv8.1-a` and builds its module with the matching feature set, so the gated LSE rows are reached on both sides. What it does not yet cover: SVE/SME/NEON (nothing to test — see above), the two gaps just listed, and the exact final bytes of a resolved cross-section reference, which needs a linker in the loop rather than a bare assembler.

## License

MIT
