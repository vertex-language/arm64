# Generates arm64's Advanced SIMD and cryptographic tables, typed helpers and
# their differential test, from one list, so the three cannot drift.
import sys, re
OUT = sys.argv[1]

ARR = {  # name: (reg const, Q, size, elem, lanes)
 "8B": ("reg.V8B",0,0,"B",8), "16B": ("reg.V16B",1,0,"B",16),
 "4H": ("reg.V4H",0,1,"H",4), "8H": ("reg.V8H",1,1,"H",8),
 "2S": ("reg.V2S",0,2,"S",2), "4S": ("reg.V4S",1,2,"S",4),
 "1D": ("reg.V1D",0,3,"D",1), "2D": ("reg.V2D",1,3,"D",2),
 "1Q": ("reg.V1Q",0,0,"Q",1),
}
BY_ELEM = {"B":["8B","16B"],"H":["4H","8H"],"S":["2S","4S"],"D":["2D"]}
WIDER = {"8B":"8H","16B":"8H","4H":"4S","8H":"4S","2S":"2D","4S":"2D"}
NARROW_OF = {"8H":"8B","4S":"4H","2D":"2S"}  # narrowing result (Q=0)
NARROW2_OF = {"8H":"16B","4S":"8H","2D":"4S"}

rows = []   # (comment-group, go row text)
methods = []
cases = []
group = None

def camel(m):
    return m[:1].upper() + m[1:]

def vslot(role, arr, field):
    return ".%sV(%s, %s)" % (role, ARR[arr][0], field)

NAMED = {"Rd":(0,5),"Rn":(5,5),"Rm":(16,5),"Ra":(10,5)}
def fields_of(slot):
    out = []
    for nm in re.findall(r"\b(Rd|Rn|Rm|Ra)\b", slot):
        out.append(NAMED[nm])
    for lo, w in re.findall(r"\bF\((\d+), (\d+)\)", slot):
        out.append((int(lo), int(w)))
    for a0, w0, a1, w1 in re.findall(r"FSplit\((\d+), (\d+), (\d+), (\d+)\)", slot):
        out += [(int(a0), int(w0)), (int(a1), int(w1))]
    return out

def add(mnem, word, mask, slots, name, gate=None, sig=None, call=None, asm=None):
    # The mask is every bit no operand writes, from the slots themselves.
    mask = 0xffffffff
    for sl in slots:
        for lo, w in fields_of(sl):
            mask &= ~(((1 << w) - 1) << lo) & 0xffffffff
    assert word & ~mask & 0xffffffff == 0, (name, hex(word), hex(mask))
    slots = [sl if sl.startswith(".") else "." + sl for sl in slots]
    gate_s = ".Gate2(feature.%s)" % gate if gate else ""
    rows.append((group, 'L("%s", %#010x, %#010x).\n\t\t\t%s%s.Name("%s"),' % (mnem, word, mask, "".join(slots)[1:], gate_s, name)))
    var = name[:1].lower() + name[1:]
    methods.append((group, var, name, sig, call))
    cases.append((name, asm, gate))

def V(i): return "v%d" % i

MASK_3SAME = 0xffe0fc00  # Rm, Rn, Rd free
MASK_2REG = 0xfffffc00

# ---- three same ----------------------------------------------------------
group = "Integer arithmetic, three registers of one arrangement"
THREE = [("add",0,0b10000,"BHSD"),("sub",1,0b10000,"BHSD"),("mul",0,0b10011,"BHS"),
 ("mla",0,0b10010,"BHS"),("mls",1,0b10010,"BHS"),
 ("cmeq",1,0b10001,"BHSD"),("cmtst",0,0b10001,"BHSD"),("cmgt",0,0b00110,"BHSD"),("cmge",0,0b00111,"BHSD"),
 ("cmhi",1,0b00110,"BHSD"),("cmhs",1,0b00111,"BHSD"),
 ("sqadd",0,0b00001,"BHSD"),("uqadd",1,0b00001,"BHSD"),("sqsub",0,0b00101,"BHSD"),("uqsub",1,0b00101,"BHSD"),
 ("smax",0,0b01100,"BHS"),("umax",1,0b01100,"BHS"),("smin",0,0b01101,"BHS"),("umin",1,0b01101,"BHS"),
 ("srhadd",0,0b00010,"BHS"),("urhadd",1,0b00010,"BHS"),("sabd",0,0b01110,"BHS"),("uabd",1,0b01110,"BHS"),("sshl",0,0b01000,"BHSD"),("ushl",1,0b01000,"BHSD"),
 ("addp",0,0b10111,"BHSD"),("sqdmulh",0,0b10110,"HS"),("sqrdmulh",1,0b10110,"HS")]
for m,u,op,elems in THREE:
    acc = m in ("mla","mls")
    for e in elems:
        for a in BY_ELEM[e]:
            _,q,size,_,_ = ARR[a]
            w = 0x0E200400 | q<<30 | u<<29 | size<<22 | op<<11
            name = camel(m)+"V"+a
            add(m, w, MASK_3SAME, [vslot("SrcDst" if acc else "Dst",a,"Rd"), vslot("Src",a,"Rn"), vslot("Src",a,"Rm")], name,
                sig="rd, rn, rm reg.V", call="rd.Arr(%s), rn.Arr(%s), rm.Arr(%s)" % ((ARR[a][0],)*3),
                asm="%s v0.%s, v1.%s, v2.%s" % (m,a.lower(),a.lower(),a.lower()))

group = "Bitwise, three registers"
for m,u,size,acc in [("and",0,0,False),("bic",0,1,False),("orr",0,2,False),("orn",0,3,False),
                      ("eor",1,0,False),("bsl",1,1,True),("bit",1,2,True),("bif",1,3,True)]:
    for a in ["8B","16B"]:
        q = ARR[a][1]
        w = 0x0E201C00 | q<<30 | u<<29 | size<<22
        name = camel(m)+"V"+a
        add(m, w, MASK_3SAME, [vslot("SrcDst" if acc else "Dst",a,"Rd"), vslot("Src",a,"Rn"), vslot("Src",a,"Rm")], name,
            sig="rd, rn, rm reg.V", call="rd.Arr(%s), rn.Arr(%s), rm.Arr(%s)" % ((ARR[a][0],)*3),
            asm="%s v0.%s, v1.%s, v2.%s" % (m,a.lower(),a.lower(),a.lower()))

# ---- two-register miscellaneous -----------------------------------------
group = "One source register"
MISC = [("rev64",0,0b00000,"BHS",None),("rev32",1,0b00000,"BH",None),("rev16",0,0b00001,"B",None),
 ("cnt",0,0b00101,"B",None),("not",1,0b00101,"B",0),("rbit",1,0b00101,"B",1),
 ("neg",1,0b01011,"BHSD",None),("abs",0,0b01011,"BHSD",None)]
for m,u,op,elems,fixsize in MISC:
    for e in elems:
        for a in BY_ELEM[e]:
            _,q,size,_,_ = ARR[a]
            if fixsize is not None: size = fixsize
            w = 0x0E200800 | q<<30 | u<<29 | size<<22 | op<<12
            name = camel(m)+"V"+a
            add(m, w, MASK_2REG, [vslot("Dst",a,"Rd"), vslot("Src",a,"Rn")], name,
                sig="rd, rn reg.V", call="rd.Arr(%s), rn.Arr(%s)" % ((ARR[a][0],)*2),
                asm="%s v0.%s, v1.%s" % (m,a.lower(),a.lower()))

group = "Narrowing and pairwise widening"
for m,u,op in [("xtn",0,0b10010),("sqxtn",0,0b10100),("uqxtn",1,0b10100),("sqxtun",1,0b10010)]:
    for wide in ["8H","4S","2D"]:
        for hi in (0,1):
            nar = NARROW2_OF[wide] if hi else NARROW_OF[wide]
            size = ARR[nar][2]
            w = 0x0E210800 | hi<<30 | u<<29 | size<<22 | op<<12
            mm = m + ("2" if hi else "")
            name = camel(m)+("2" if hi else "")+"V"+nar
            add(mm, w, MASK_2REG, [vslot("SrcDst" if hi else "Dst",nar,"Rd"), vslot("Src",wide,"Rn")], name,
                sig="rd, rn reg.V", call="rd.Arr(%s), rn.Arr(%s)" % (ARR[nar][0], ARR[wide][0]),
                asm="%s v0.%s, v1.%s" % (mm, nar.lower(), wide.lower()))
for m,u in [("uaddlp",1),("saddlp",0)]:
    for src,dst in [("8B","4H"),("16B","8H"),("4H","2S"),("8H","4S"),("2S","1D"),("4S","2D")]:
        _,q,size,_,_ = ARR[src]
        w = 0x0E202800 | q<<30 | u<<29 | size<<22
        name = camel(m)+"V"+dst
        add(m, w, MASK_2REG, [vslot("Dst",dst,"Rd"), vslot("Src",src,"Rn")], name,
            sig="rd, rn reg.V", call="rd.Arr(%s), rn.Arr(%s)" % (ARR[dst][0], ARR[src][0]),
            asm="%s v0.%s, v1.%s" % (m, dst.lower(), src.lower()))

group = "Across all lanes, into a scalar"
SC = {"B":("ClassB","reg.B","b"),"H":("ClassH","reg.H","h"),"S":("ClassS","reg.S","s"),"D":("ClassD","reg.D","d")}
UP = {"B":"H","H":"S","S":"D"}
for m,u,op,long in [("addv",0,0b11011,False),("umaxv",1,0b01010,False),("uminv",1,0b11010,False),
                    ("smaxv",0,0b01010,False),("sminv",0,0b11010,False),("uaddlv",1,0b00011,True),("saddlv",0,0b00011,True)]:
    for a in ["8B","16B","4H","8H","4S"]:
        _,q,size,e,_ = ARR[a]
        de = UP[e] if long else e
        cls, gotype, letter = SC[de]
        w = 0x0E300800 | q<<30 | u<<29 | size<<22 | op<<12
        name = camel(m)+"V"+a
        add(m, w, MASK_2REG, ["Dst(%s, Rd)" % cls, vslot("Src",a,"Rn")], name,
            sig="rd %s, rn reg.V" % gotype, call="rd, rn.Arr(%s)" % ARR[a][0],
            asm="%s %s0, v1.%s" % (m, letter, a.lower()))

# ---- shift by immediate -------------------------------------------------
group = "Shifts by an immediate"
SH = [("shl",0,0b01010,"L",False),("sli",1,0b01010,"L",True),("ushr",1,0b00000,"R",False),("sshr",0,0b00000,"R",False),
      ("usra",1,0b00010,"R",True),("ssra",0,0b00010,"R",True),("urshr",1,0b00100,"R",False),("srshr",0,0b00100,"R",False),
      ("sri",1,0b01000,"R",True)]
for m,u,op,dirn,acc in SH:
    for e in "BHSD":
        for a in BY_ELEM[e]:
            q = ARR[a][1]
            w = 0x0F000400 | q<<30 | u<<29 | op<<11
            name = camel(m)+"V"+a
            kind = "ImmVecShl" if dirn=="L" else "ImmVecShr"
            amt = 3 if dirn=="L" else 5
            add(m, w, 0xff80fc00, [vslot("SrcDst" if acc else "Dst",a,"Rd"), vslot("Src",a,"Rn"), ".Imm(F(16, 7)).Kind(%s)" % kind], name,
                sig="rd, rn reg.V, shift uint8", call="rd.Arr(%s), rn.Arr(%s), uint64(shift)" % ((ARR[a][0],)*2),
                asm="%s v0.%s, v1.%s, #%d" % (m,a.lower(),a.lower(),amt))
for m,u,op in [("shrn",0,0b10000),("rshrn",0,0b10001),("sqshrn",0,0b10010),("uqshrn",1,0b10010),("sqshrun",1,0b10000)]:
    for wide in ["8H","4S","2D"]:
        for hi in (0,1):
            nar = NARROW2_OF[wide] if hi else NARROW_OF[wide]
            w = 0x0F000400 | hi<<30 | u<<29 | op<<11
            mm = m + ("2" if hi else "")
            name = camel(m)+("2" if hi else "")+"V"+nar
            add(mm, w, 0xff80fc00, [vslot("SrcDst" if hi else "Dst",nar,"Rd"), vslot("Src",wide,"Rn"), ".Imm(F(16, 7)).Kind(ImmVecShr)"], name,
                sig="rd, rn reg.V, shift uint8", call="rd.Arr(%s), rn.Arr(%s), uint64(shift)" % (ARR[nar][0], ARR[wide][0]),
                asm="%s v0.%s, v1.%s, #3" % (mm, nar.lower(), wide.lower()))
for m,u in [("ushll",1),("sshll",0)]:
    for src in ["8B","16B","4H","8H","2S","4S"]:
        q = ARR[src][1]
        dst = WIDER[src]
        w = 0x0F00A400 | q<<30 | u<<29
        mm = m + ("2" if q else "")
        name = camel(m)+("2" if q else "")+"V"+dst
        add(mm, w, 0xff80fc00, [vslot("Dst",dst,"Rd"), vslot("Src",src,"Rn"), ".Imm(F(16, 7)).Kind(ImmVecShl)"], name,
            sig="rd, rn reg.V, shift uint8", call="rd.Arr(%s), rn.Arr(%s), uint64(shift)" % (ARR[dst][0], ARR[src][0]),
            asm="%s v0.%s, v1.%s, #2" % (mm, dst.lower(), src.lower()))

# ---- three different --------------------------------------------------
group = "Widening, three registers"
for m,u,op,acc in [("umull",1,0b1100,False),("smull",0,0b1100,False),("umlal",1,0b1000,True),("smlal",0,0b1000,True),
                   ("uaddl",1,0b0000,False),("saddl",0,0b0000,False),("usubl",1,0b0010,False),("ssubl",0,0b0010,False)]:
    for src in ["8B","16B","4H","8H","2S","4S"]:
        _,q,size,_,_ = ARR[src]
        dst = WIDER[src]
        w = 0x0E200000 | q<<30 | u<<29 | size<<22 | op<<12
        mm = m + ("2" if q else "")
        name = camel(m)+("2" if q else "")+"V"+dst
        add(mm, w, MASK_3SAME, [vslot("SrcDst" if acc else "Dst",dst,"Rd"), vslot("Src",src,"Rn"), vslot("Src",src,"Rm")], name,
            sig="rd, rn, rm reg.V", call="rd.Arr(%s), rn.Arr(%s), rm.Arr(%s)" % (ARR[dst][0], ARR[src][0], ARR[src][0]),
            asm="%s v0.%s, v1.%s, v2.%s" % (mm, dst.lower(), src.lower(), src.lower()))
for src,dst,q,size,gate in [("8B","8H",0,0,None),("16B","8H",1,0,None),("1D","1Q",0,3,"AES"),("2D","1Q",1,3,"AES")]:
    w = 0x0E20E000 | q<<30 | size<<22
    mm = "pmull" + ("2" if q else "")
    name = "Pmull"+("2" if q else "")+"V"+dst
    add(mm, w, MASK_3SAME, [vslot("Dst",dst,"Rd"), vslot("Src",src,"Rn"), vslot("Src",src,"Rm")], name, gate=gate,
        sig="rd, rn, rm reg.V", call="rd.Arr(%s), rn.Arr(%s), rm.Arr(%s)" % (ARR[dst][0], ARR[src][0], ARR[src][0]),
        asm="%s v0.%s, v1.%s, v2.%s" % (mm, dst.lower(), src.lower(), src.lower()))

# ---- permutes ------------------------------------------------------
group = "Permutes"
for m,op in [("uzp1",1),("trn1",2),("zip1",3),("uzp2",5),("trn2",6),("zip2",7)]:
    for a in ["8B","16B","4H","8H","2S","4S","2D"]:
        _,q,size,_,_ = ARR[a]
        w = 0x0E000800 | q<<30 | size<<22 | op<<12
        name = camel(m)+"V"+a
        add(m, w, MASK_3SAME, [vslot("Dst",a,"Rd"), vslot("Src",a,"Rn"), vslot("Src",a,"Rm")], name,
            sig="rd, rn, rm reg.V", call="rd.Arr(%s), rn.Arr(%s), rm.Arr(%s)" % ((ARR[a][0],)*3),
            asm="%s v0.%s, v1.%s, v2.%s" % (m,a.lower(),a.lower(),a.lower()))
for a,q in [("8B",0),("16B",1)]:
    w = 0x2E000000 | q<<30
    add("ext", w, 0xffe08400, [vslot("Dst",a,"Rd"), vslot("Src",a,"Rn"), vslot("Src",a,"Rm"), ".Imm(F(11, 4))"], "ExtV"+a,
        sig="rd, rn, rm reg.V, index uint8", call="rd.Arr(%s), rn.Arr(%s), rm.Arr(%s), uint64(index)" % ((ARR[a][0],)*3),
        asm="ext v0.%s, v1.%s, v2.%s, #%d" % (a.lower(),a.lower(),a.lower(), 3 if q==0 else 9))
# TBL and TBX of one table register: the text spells the table {v1.16b}.
for m,bit,acc in [("tbl",0,False),("tbx",1,True)]:
    for a,q in [("8B",0),("16B",1)]:
        w = 0x0E000000 | q<<30 | bit<<12
        add(m, w, MASK_3SAME, [vslot("SrcDst" if acc else "Dst",a,"Rd"), vslot("Src","16B","Rn"), vslot("Src",a,"Rm")], camel(m)+"V"+a,
            sig="rd, table, index reg.V", call="rd.Arr(%s), table.Arr(reg.V16B), index.Arr(%s)" % (ARR[a][0], ARR[a][0]),
            asm="%s v0.%s, {v1.16b}, v2.%s" % (m, a.lower(), a.lower()))

# ---- moves between lanes and general registers -----------------------
group = "Lanes and general registers"
IMM5 = {"B":(1,F(17,4) if False else "F(17, 4)"),"H":(2,"F(18, 3)"),"S":(4,"F(19, 2)"),"D":(8,"F(20, 1)")}
ELEMC = {"B":"reg.ElemB","H":"reg.ElemH","S":"reg.ElemS","D":"reg.ElemD"}
for e in "BHSD":
    pin, idx = IMM5[e]
    for a in BY_ELEM[e]:
        q = ARR[a][1]
        gpr = ("ClassX","reg.X","x") if e=="D" else ("ClassW","reg.W","w")
        w = 0x0E000C00 | q<<30 | pin<<16
        add("dup", w, MASK_2REG, [vslot("Dst",a,"Rd"), "Src(%s, Rn)" % gpr[0]], "DupV%sGen" % a,
            sig="rd reg.V, rn %s" % gpr[1], call="rd.Arr(%s), rn" % ARR[a][0],
            asm="dup v0.%s, %s1" % (a.lower(), gpr[2]))
        w = 0x0E000400 | q<<30 | pin<<16
        add("dup", w, 0xffe0fc00 | (0) , [vslot("Dst",a,"Rd"), ".SrcLane(%s, Rn, %s)" % (ELEMC[e], idx)], "DupV%sLane" % a,
            sig="rd, rn reg.V, index uint8", call="rd.Arr(%s), rn.Lane(%s, index)" % (ARR[a][0], ELEMC[e]),
            asm="dup v0.%s, v1.%s[1]" % (a.lower(), e.lower()))
    # UMOV and INS.
    gpr = ("ClassX","reg.X","x") if e=="D" else ("ClassW","reg.W","w")
    q = 1 if e=="D" else 0
    w = 0x0E003C00 | q<<30 | pin<<16
    add("umov", w, 0xffe0fc00, ["Dst(%s, Rd)" % gpr[0], ".SrcLane(%s, Rn, %s)" % (ELEMC[e], idx)], "Umov%s" % e,
        sig="rd %s, rn reg.V, index uint8" % gpr[1], call="rd, rn.Lane(%s, index)" % ELEMC[e],
        asm="umov %s0, v1.%s[1]" % (gpr[2], e.lower()))
    w = 0x4E001C00 | pin<<16
    add("ins", w, 0xffe0fc00, [".DstLane(%s, Rd, %s)" % (ELEMC[e], idx), "Src(%s, Rn)" % gpr[0]], "Ins%sGen" % e,
        sig="rd reg.V, index uint8, rn %s" % gpr[1], call="rd.Lane(%s, index), rn" % ELEMC[e],
        asm="ins v0.%s[1], %s1" % (e.lower(), gpr[2]))
for e,pin,idx,q in [("B",1,"F(17, 4)",0),("H",2,"F(18, 3)",0),("B",1,"F(17, 4)",1),("H",2,"F(18, 3)",1),("S",4,"F(19, 2)",1)]:
    gpr = ("ClassX","reg.X","x") if q else ("ClassW","reg.W","w")
    w = 0x0E002C00 | q<<30 | pin<<16
    add("smov", w, 0xffe0fc00, ["Dst(%s, Rd)" % gpr[0], ".SrcLane(%s, Rn, %s)" % (ELEMC[e], idx)], "Smov%s%s" % (e, "64" if q else "32"),
        sig="rd %s, rn reg.V, index uint8" % gpr[1], call="rd, rn.Lane(%s, index)" % ELEMC[e],
        asm="smov %s0, v1.%s[1]" % (gpr[2], e.lower()))

group = "Immediates into a vector"
add("movi", 0x6F00E400, 0xfff8fc00, [vslot("Dst","2D","Rd"), ".Imm(FSplit(5, 5, 16, 3)).Kind(ImmVecByteMask)"], "MoviV2D",
    sig="rd reg.V, mask uint64", call="rd.Arr(reg.V2D), mask", asm="movi v0.2d, #0xff00ff0000ff00ff")
for a,q in [("8B",0),("16B",1)]:
    add("movi", 0x0F00E400 | q<<30, 0xfff8fc00, [vslot("Dst",a,"Rd"), ".Imm(FSplit(5, 5, 16, 3))"], "MoviV"+a,
        sig="rd reg.V, imm uint8", call="rd.Arr(%s), uint64(imm)" % ARR[a][0], asm="movi v0.%s, #0xa5" % a.lower())

# ---- cryptography ----------------------------------------------------
group = "AES"
for m,w in [("aese",0x4E284800),("aesd",0x4E285800)]:
    add(m, w, MASK_2REG, [vslot("SrcDst","16B","Rd"), vslot("Src","16B","Rn")], camel(m), gate="AES",
        sig="rd, rn reg.V", call="rd.Arr(reg.V16B), rn.Arr(reg.V16B)", asm="%s v0.16b, v1.16b" % m)
for m,w in [("aesmc",0x4E286800),("aesimc",0x4E287800)]:
    add(m, w, MASK_2REG, [vslot("Dst","16B","Rd"), vslot("Src","16B","Rn")], camel(m), gate="AES",
        sig="rd, rn reg.V", call="rd.Arr(reg.V16B), rn.Arr(reg.V16B)", asm="%s v0.16b, v1.16b" % m)

group = "SHA-1 and SHA-256"
for m,w in [("sha1c",0x5E000000),("sha1p",0x5E001000),("sha1m",0x5E002000)]:
    add(m, w, MASK_3SAME, ["SrcDst(ClassQ, Rd)", "Src(ClassS, Rn)", vslot("Src","4S","Rm")], camel(m), gate="SHA2",
        sig="rd reg.Q, rn reg.S, rm reg.V", call="rd, rn, rm.Arr(reg.V4S)", asm="%s q0, s1, v2.4s" % m)
add("sha1su0", 0x5E003000, MASK_3SAME, [vslot("SrcDst","4S","Rd"), vslot("Src","4S","Rn"), vslot("Src","4S","Rm")], "Sha1su0", gate="SHA2",
    sig="rd, rn, rm reg.V", call="rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S)", asm="sha1su0 v0.4s, v1.4s, v2.4s")
for m,w in [("sha256h",0x5E004000),("sha256h2",0x5E005000)]:
    add(m, w, MASK_3SAME, ["SrcDst(ClassQ, Rd)", "Src(ClassQ, Rn)", vslot("Src","4S","Rm")], camel(m), gate="SHA2",
        sig="rd, rn reg.Q, rm reg.V", call="rd, rn, rm.Arr(reg.V4S)", asm="%s q0, q1, v2.4s" % m)
add("sha256su1", 0x5E006000, MASK_3SAME, [vslot("SrcDst","4S","Rd"), vslot("Src","4S","Rn"), vslot("Src","4S","Rm")], "Sha256su1", gate="SHA2",
    sig="rd, rn, rm reg.V", call="rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S)", asm="sha256su1 v0.4s, v1.4s, v2.4s")
add("sha1h", 0x5E280800, MASK_2REG, ["Dst(ClassS, Rd)", "Src(ClassS, Rn)"], "Sha1h", gate="SHA2",
    sig="rd, rn reg.S", call="rd, rn", asm="sha1h s0, s1")
for m,w in [("sha1su1",0x5E281800),("sha256su0",0x5E282800)]:
    add(m, w, MASK_2REG, [vslot("SrcDst","4S","Rd"), vslot("Src","4S","Rn")], camel(m), gate="SHA2",
        sig="rd, rn reg.V", call="rd.Arr(reg.V4S), rn.Arr(reg.V4S)", asm="%s v0.4s, v1.4s" % m)

group = "SHA-512 and SHA-3"
for m,w in [("sha512h",0xCE608000),("sha512h2",0xCE608400)]:
    add(m, w, MASK_3SAME, ["SrcDst(ClassQ, Rd)", "Src(ClassQ, Rn)", vslot("Src","2D","Rm")], camel(m), gate="SHA3",
        sig="rd, rn reg.Q, rm reg.V", call="rd, rn, rm.Arr(reg.V2D)", asm="%s q0, q1, v2.2d" % m)
add("sha512su1", 0xCE608800, MASK_3SAME, [vslot("SrcDst","2D","Rd"), vslot("Src","2D","Rn"), vslot("Src","2D","Rm")], "Sha512su1", gate="SHA3",
    sig="rd, rn, rm reg.V", call="rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D)", asm="sha512su1 v0.2d, v1.2d, v2.2d")
add("sha512su0", 0xCEC08000, MASK_2REG, [vslot("SrcDst","2D","Rd"), vslot("Src","2D","Rn")], "Sha512su0", gate="SHA3",
    sig="rd, rn reg.V", call="rd.Arr(reg.V2D), rn.Arr(reg.V2D)", asm="sha512su0 v0.2d, v1.2d")
for m,w in [("eor3",0xCE000000),("bcax",0xCE200000)]:
    add(m, w, 0xffe08000, [vslot("Dst","16B","Rd"), vslot("Src","16B","Rn"), vslot("Src","16B","Rm"), vslot("Src","16B","Ra")], camel(m), gate="SHA3",
        sig="rd, rn, rm, ra reg.V", call="rd.Arr(reg.V16B), rn.Arr(reg.V16B), rm.Arr(reg.V16B), ra.Arr(reg.V16B)", asm="%s v0.16b, v1.16b, v2.16b, v3.16b" % m)
add("rax1", 0xCE608C00, MASK_3SAME, [vslot("Dst","2D","Rd"), vslot("Src","2D","Rn"), vslot("Src","2D","Rm")], "Rax1", gate="SHA3",
    sig="rd, rn, rm reg.V", call="rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D)", asm="rax1 v0.2d, v1.2d, v2.2d")
add("xar", 0xCE800000, 0xffe00000, [vslot("Dst","2D","Rd"), vslot("Src","2D","Rn"), vslot("Src","2D","Rm"), ".Imm(F(10, 6))"], "Xar", gate="SHA3",
    sig="rd, rn, rm reg.V, rotate uint8", call="rd.Arr(reg.V2D), rn.Arr(reg.V2D), rm.Arr(reg.V2D), uint64(rotate)", asm="xar v0.2d, v1.2d, v2.2d, #13")

group = "SM3 and SM4"
add("sm3ss1", 0xCE400000, 0xffe08000, [vslot("Dst","4S","Rd"), vslot("Src","4S","Rn"), vslot("Src","4S","Rm"), vslot("Src","4S","Ra")], "Sm3ss1", gate="SM4",
    sig="rd, rn, rm, ra reg.V", call="rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S), ra.Arr(reg.V4S)", asm="sm3ss1 v0.4s, v1.4s, v2.4s, v3.4s")
for m,w in [("sm3tt1a",0xCE408000),("sm3tt1b",0xCE408400),("sm3tt2a",0xCE408800),("sm3tt2b",0xCE408C00)]:
    add(m, w, 0xffe0cc00, [vslot("SrcDst","4S","Rd"), vslot("Src","4S","Rn"), ".SrcLane(reg.ElemS, Rm, F(12, 2))"], camel(m), gate="SM4",
        sig="rd, rn, rm reg.V, index uint8", call="rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Lane(reg.ElemS, index)", asm="%s v0.4s, v1.4s, v2.s[2]" % m)
for m,w in [("sm3partw1",0xCE60C000),("sm3partw2",0xCE60C400),("sm4ekey",0xCE60C800)]:
    role = "Dst" if m=="sm4ekey" else "SrcDst"
    add(m, w, MASK_3SAME, [vslot(role,"4S","Rd"), vslot("Src","4S","Rn"), vslot("Src","4S","Rm")], camel(m), gate="SM4",
        sig="rd, rn, rm reg.V", call="rd.Arr(reg.V4S), rn.Arr(reg.V4S), rm.Arr(reg.V4S)", asm="%s v0.4s, v1.4s, v2.4s" % m)
add("sm4e", 0xCEC08400, MASK_2REG, [vslot("SrcDst","4S","Rd"), vslot("Src","4S","Rn")], "Sm4e", gate="SM4",
    sig="rd, rn reg.V", call="rd.Arr(reg.V4S), rn.Arr(reg.V4S)", asm="sm4e v0.4s, v1.4s")

group = "CRC-32 and CRC-32C"
for m,w,x in [("crc32b",0x1AC04000,False),("crc32h",0x1AC04400,False),("crc32w",0x1AC04800,False),("crc32x",0x9AC04C00,True),
              ("crc32cb",0x1AC05000,False),("crc32ch",0x1AC05400,False),("crc32cw",0x1AC05800,False),("crc32cx",0x9AC05C00,True)]:
    add(m, w, MASK_3SAME, ["Dst(ClassW, Rd)", "Src(ClassW, Rn)", "Src(%s, Rm)" % ("ClassX" if x else "ClassW")], camel(m), gate="CRC",
        sig="rd, rn reg.W, rm %s" % ("reg.X" if x else "reg.W"), call="rd, rn, rm", asm="%s w0, w1, %s2" % (m, "x" if x else "w"))

# ---- emit -------------------------------------------------------------
def fix(slot):
    slot = slot.strip()
    return slot if slot.startswith(".") else "." + slot
tbl = ['''package isa

import (
	"github.com/vertex-language/arm64/feature"
	"github.com/vertex-language/arm64/reg"
)

// Advanced SIMD's integer instructions and the cryptographic extensions.
// Generated from one list with inst_simd.go and simd_difftest_test.go, so
// a row, its typed helper and the clang case that checks its word cannot
// disagree; edit the list in gen_simd.py, not the output, and regenerate
// with "python3 internal/isa/gen_simd.py /tmp/out" and copying the three
// files into place.
//
// A vector operand names its arrangement: the size and Q bits each row
// pins are what tell v0.4s from v0.8h, so a slot takes the one its row
// encodes. Advanced SIMD is in the baseline; the cryptographic rows are
// gated on their features, PMULL's 64-bit form on AES as the architecture
// has it.
func init() {
	register(''']
last = None
for g, r in rows:
    if g != last:
        tbl.append("\n\t\t// ---- %s ----" % g)
        last = g
    # normalize slot method chains
    head, rest = r.split("\n", 1)
    rest = rest.strip()
    tbl.append("\t\t" + head + "\n\t\t\t" + rest)
tbl.append("\t)\n}\n")
src = "\n".join(tbl)
# slots were written as ".DstV(...)" or "Dst(...)"; make every slot a chained call
import re
open(OUT + "/internal/isa/table_simd.go", "w").write(src)

inst = ['''package arm64

import "github.com/vertex-language/arm64/reg"

// Advanced SIMD and the cryptographic extensions, one helper per row of
// internal/isa/table_simd.go: the arrangement is in the name -- AddV4S adds
// four 32-bit lanes -- and the registers are plain V, decorated here. A
// lane is a register and an index. Generated with the table; see there.
''']
last = None
for g, var, name, sig, call in methods:
    if g != last:
        inst.append("\n// ---- %s ----\n" % g)
        last = g
    inst.append("var %s = form(%q)" % (var, name) if False else 'var %s = form("%s")' % (var, name))
    inst.append("\nfunc (s *Section) %s(%s) { s.inst(%s, %s) }\n" % (name, sig, var, call))
open(OUT + "/inst_simd.go", "w").write("\n".join(inst))

import re
PREFIX = {"reg.V":"V","reg.W":"W","reg.X":"X","reg.Q":"Q","reg.S":"S","reg.B":"B","reg.H":"H","reg.D":"D"}
def tcall(name, sig, asm):
    # sig "rd, rn, rm reg.V, shift uint8" -> [(name, type)]
    params = []
    pending = []
    for part in sig.split(","):
        bits = part.strip().split()
        pending.append(bits[0])
        if len(bits) == 2:
            for n in pending:
                params.append((n, bits[1]))
            pending = []
    nums = re.findall(r"#(0x[0-9a-f]+|\d+)", asm)
    lane = re.findall(r"\[(\d+)\]", asm)
    out, regi = [], 0
    for n, t in params:
        if t in PREFIX:
            out.append("arm64.%s%d" % (PREFIX[t], regi))
            regi += 1
        elif n == "index" and lane:
            out.append(lane[0])
        else:
            out.append(nums.pop(0))
    return "t.%s(%s)" % (name, ", ".join(out))

test = ["package arm64_test",
"",
"// The Advanced SIMD and cryptographic rows against clang's assembler, one",
"// case per row: the differential check difftest_test.go makes of the base",
"// set. Generated with internal/isa/table_simd.go; edit the list there.",
"",
"import (",
"\t\"bytes\"",
"\t\"encoding/binary\"",
"\t\"os\"",
"\t\"os/exec\"",
"\t\"path/filepath\"",
"\t\"strings\"",
"\t\"testing\"",
"",
"\t\"github.com/vertex-language/arm64\"",
"\t\"github.com/vertex-language/arm64/feature\"",
")",
"",
"func simdCases() []diffCase {",
"\treturn []diffCase{"]
for name, asm, gate in cases:
    sig = [x for x in methods if x[2]==name][0][3]
    test.append('\t\t{"%s", "%s", func(t *arm64.Section) { %s }},' % (name, asm, tcall(name, sig, asm)))
test.append("""\t}
}

func TestSIMDAgainstClang(t *testing.T) {
	clangPath, err := exec.LookPath("clang")
	if err != nil {
		t.Skip("clang not on PATH; skipping differential test against the reference assembler")
	}
	cases := simdCases()
	var src strings.Builder
	src.WriteString(".arch armv8.4-a+crypto+sha3+sm4\\n")
	for _, c := range cases {
		src.WriteString(c.asm)
		src.WriteByte('\\n')
	}
	dir := t.TempDir()
	srcPath := filepath.Join(dir, "simd.s")
	objPath := filepath.Join(dir, "simd.o")
	if err := os.WriteFile(srcPath, []byte(src.String()), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(clangPath, "-target", "aarch64-linux-gnu", "-c", "-o", objPath, srcPath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("clang: %v\\n%s", err, stderr.String())
	}
	want, err := textWords(objPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(want) != len(cases) {
		t.Fatalf("clang produced %d words for %d cases", len(want), len(cases))
	}
	m := arm64.NewModule(arm64.WithFeatures(arm64.Armv8_4A.Set().Plus(feature.AES, feature.SHA2, feature.SHA3, feature.SM4)))
	sec := m.Section(arm64.Text)
	for _, c := range cases {
		c.build(sec)
	}
	if err := m.Err(); err != nil {
		t.Fatalf("building the typed-helper sequence: %v", err)
	}
	obj, err := m.Finalize()
	if err != nil {
		t.Fatalf("Finalize: %v", err)
	}
	got := obj.SectionAt(0).Bytes()
	if len(got) != len(cases)*4 {
		t.Fatalf("got %d bytes for %d cases", len(got), len(cases))
	}
	for i, c := range cases {
		if g := binary.LittleEndian.Uint32(got[i*4:]); g != want[i] {
			t.Errorf("%-22s %-40s got %#08x, want %#08x (clang)", c.name, c.asm, g, want[i])
		}
	}
}
""")
open(OUT + "/simd_difftest_test.go", "w").write("\n".join(test))
print(len(rows), "rows")
