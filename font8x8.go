//    Original Source - Fetched July 3nd 2014
//      http://www.d-rift.nl/combuster/mos3/?p=viewsource&file=/modules/gfx/font8_8.asm
//
//    Author:  Marcel Sondaar
//             International Business Machines (public domain VGA fonts)
//    License: Public Domain
//
//    Converted July 3rd 2014 By Jeremy Jay for use with github.com/pbnjay/pixfont
//      -- Removed 781 empty and "squiggle-block" 442A112A442A1100 glyphs
//

package pixfont

// unexported to make godoc cleaner...
var eightMap = map[int32]uint16{9607: 0x2, 20114: 0x2a0, 9519: 0x1cc2, 9515: 0x2402, 20103: 0x2760, 12380: 0x840, 12407: 0x15a0, 203: 0x1680, 24037: 0x1a62, 9563: 0x2920, 9566: 0x2880, 9526: 0x680, 63: 0x942, 9569: 0xe40, 54: 0xf62, 21316: 0x15a2, 9587: 0x20e2, 945: 0x122, 20116: 0xac2, 957: 0x14a2, 12374: 0x14e0, 12403: 0x14e2, 19970: 0x19e2, 12357: 0x2080, 20051: 0x1140, 58697: 0x1480, 9555: 0x14c2, 9571: 0x1602, 220: 0x1880, 249: 0x20e0, 9481: 0x2a80, 23667: 0x9a2, 21313: 0xf42, 225: 0x1582, 9557: 0x1ce0, 19972: 0x1d60, 104: 0xd02, 12375: 0x1ba2, 20122: 0x2ae2, 58711: 0x360, 20010: 0x520, 58698: 0xb60, 12382: 0xbe0, 9579: 0xce0, 20837: 0xee0, 223: 0x2242, 9539: 0x2522, 12426: 0x2802, 22235: 0xfe0, 58695: 0x1342, 21494: 0x1642, 9484: 0x17a0, 20844: 0x2020, 941: 0x12c2, 194: 0x1580, 22236: 0x1740, 175: 0x262, 12429: 0xaa0, 20846: 0xe60, 947: 0x1002, 924: 0x11a2, 9494: 0x2800, 207: 0x422, 943: 0xd22, 12376: 0x24c2, 9562: 0x2600, 9550: 0x26e0, 12425: 0x5e2, 21317: 0xdc2, 12444: 0x17a2, 23668: 0x1962, 116: 0x29c0, 12410: 0xfa2, 19971: 0x15e0, 238: 0x1760, 9533: 0x19c0, 196: 0x20c2, 9626: 0x2102, 58696: 0x3c0, 9608: 0xf22, 125: 0x1322, 9493: 0x1702, 9490: 0x1c00, 57: 0x1162, 246: 0x1be2, 22233: 0x20a2, 24029: 0x28e0, 12390: 0x1c02, 9611: 0x2300, 74: 0x25a2, 192: 0x60, 23670: 0x100, 20015: 0x8c0, 21498: 0x1502, 9485: 0x19e0, 12443: 0x2782, 9560: 0x2382, 172: 0x27e0, 20003: 0x1d00, 9529: 0x1ea0, 9542: 0xa0, 31348: 0x642, 9473: 0x722, 9530: 0x982, 58689: 0xde2, 102: 0x15c0, 12446: 0x16c2, 9501: 0x1900, 84: 0x420, 12402: 0x442, 12396: 0x4a2, 919: 0xe20, 21490: 0x1282, 30001: 0x1ec2, 83: 0x1fe2, 94: 0x24e2, 217: 0x1f80, 12378: 0x2162, 912: 0x2320, 950: 0x3c2, 9520: 0xc40, 8976: 0x1982, 56: 0x1a80, 9517: 0x1f22, 12428: 0x2562, 33463: 0x2360, 185: 0x2460, 12411: 0x2aa2, 20062: 0x282, 20838: 0x1aa2, 42: 0x1b40, 58694: 0x1d20, 97: 0x21a0, 80: 0x2902, 202: 0xe00, 9604: 0x1720, 68: 0x21e0, 9617: 0x2340, 925: 0x24a2, 9597: 0x920, 9593: 0xf20, 26093: 0x1800, 101: 0x1fe0, 961: 0x1d42, 168: 0x22e2, 9577: 0x2560, 9488: 0x140, 9631: 0xa82, 936: 0x1362, 100: 0x1562, 22238: 0x1742, 9600: 0x1ae2, 24191: 0x1e42, 9612: 0x2822, 12355: 0x80, 12358: 0xc80, 9537: 0x1100, 58702: 0x13a0, 31349: 0x1660, 113: 0x320, 12405: 0xb82, 22231: 0xd60, 177: 0x1080, 9589: 0x22c0, 20008: 0x23e0, 108: 0x29a2, 9510: 0x9e0, 103: 0x11a0, 37: 0x16a2, 95: 0x18c2, 20847: 0x1ac0, 927: 0x2000, 968: 0x21c2, 9524: 0x2322, 9610: 0x2680, 58714: 0x13a2, 30004: 0x1600, 9546: 0x1a22, 245: 0x1e2, 19979: 0x202, 22825: 0x4e0, 9605: 0xea0, 9538: 0x10e2, 914: 0x1f20, 20057: 0x2420, 40: 0x2440, 12431: 0x2520, 19975: 0x2700, 55: 0x82, 26411: 0x402, 21314: 0x462, 62: 0xa20, 23567: 0x1d02, 22823: 0x16e2, 12427: 0x1c82, 12361: 0x26e2, 85: 0x3e0, 22303: 0x9a0, 9487: 0xa60, 9568: 0xbc2, 969: 0x16c0, 9551: 0x820, 962: 0x2160, 9603: 0x23a0, 9513: 0x2a82, 12399: 0x6c2, 215: 0x1400, 12435: 0x1940, 935: 0x2ac2, 58705: 0x2702, 247: 0x742, 20839: 0x1f40, 934: 0x2002, 937: 0x21c0, 26412: 0x24e0, 9575: 0x2500, 958: 0x3e2, 213: 0x482, 9541: 0x1de2, 19969: 0x1fc0, 81: 0x24a0, 9514: 0x25e2, 9578: 0x2642, 35: 0x8e0, 19973: 0x960, 20124: 0xa02, 30008: 0xd00, 49: 0x1062, 9545: 0x2740, 60: 0x260, 20009: 0x322, 58: 0x500, 12397: 0x1380, 920: 0x1fa0, 930: 0x2a00, 953: 0x380, 122: 0x480, 12386: 0xd40, 12423: 0x1842, 22829: 0x1ec0, 949: 0x1262, 952: 0x1280, 12381: 0x1ee2, 9570: 0x142, 9483: 0x900, 226: 0xa42, 948: 0xac0, 9624: 0xea2, 944: 0x280, 24180: 0x17e2, 9503: 0x25e0, 26413: 0x522, 12383: 0xbe2, 12415: 0x1922, 9509: 0x2960, 20013: 0x842, 960: 0x11c0, 164: 0x1460, 183: 0x2442, 12419: 0x740, 24027: 0x1980, 9576: 0x1a40, 12420: 0x1aa0, 8359: 0x1ca0, 112: 0x1d40, 26379: 0x1e60, 942: 0x26c2, 41: 0x242, 9573: 0x902, 12395: 0x18a2, 9618: 0x1a42, 19977: 0x1b00, 30005: 0x2c2, 26092: 0x600, 9619: 0x13e0, 9479: 0x14c0, 9547: 0x1a20, 21325: 0x2942, 20104: 0x222, 12384: 0x362, 20112: 0xdc0, 20840: 0x1fc2, 26089: 0x2262, 29983: 0x582, 916: 0x8a0, 38: 0x1500, 12385: 0x1f00, 121: 0x2082, 228: 0xb20, 163: 0x1e62, 9629: 0x2602, 9616: 0x28c0, 9594: 0x2a22, 9523: 0x2780, 9591: 0x1020, 21493: 0x12a0, 26087: 0x1b20, 26414: 0x1dc2, 20119: 0x2580, 36: 0x2140, 26376: 0x700, 64: 0x800, 9521: 0x8a2, 20845: 0x1182, 19978: 0x1300, 58699: 0x2182, 12366: 0x2502, 22824: 0x962, 24179: 0xd62, 12389: 0xec0, 70: 0x14a0, 243: 0x1f82, 12368: 0x1e00, 12398: 0x1e22, 24183: 0x2480, 12369: 0xa2, 9623: 0x1120, 12417: 0x1122, 9622: 0x1c62, 9527: 0x1da0, 20006: 0x3a2, 9614: 0x1022, 236: 0x1662, 9532: 0x2542, 235: 0x760, 26408: 0x18c0, 12388: 0x2342, 26081: 0x862, 940: 0xd82, 9558: 0x1b42, 19982: 0x20a0, 22826: 0x1a2, 58690: 0x1382, 20126: 0x1542, 106: 0x1d22, 30000: 0x25a0, 12363: 0xe42, 22239: 0xfe2, 12436: 0x1340, 9497: 0x24c0, 9507: 0x1682, 9498: 0x2722, 9615: 0x28a2, 24028: 0xc2, 90: 0xae2, 12400: 0xb80, 119: 0xfc0, 66: 0x1220, 9585: 0xda2, 12371: 0x1c80, 24039: 0x1e40, 9478: 0x2a02, 9592: 0x2e0, 9508: 0x5c2, 9621: 0x440, 252: 0xf02, 79: 0x1d80, 20108: 0x8c2, 22237: 0x1142, 22830: 0x1960, 107: 0x2842, 9606: 0x2922, 77: 0xe80, 19976: 0x1202, 187: 0x1e20, 20843: 0x340, 12408: 0x562, 9472: 0x720, 12394: 0xc0, 211: 0xe2, 91: 0x180, 188: 0x1e0, 951: 0x200, 917: 0x13c2, 9581: 0x1622, 12373: 0x1a82, 9543: 0x880, 923: 0x1560, 118: 0x25c0, 26415: 0x580, 23664: 0xf60, 47: 0x1000, 21497: 0x1040, 20102: 0x1b02, 190: 0x300, 12367: 0x620, 222: 0xb02, 12445: 0x460, 182: 0x5a2, 45: 0xd80, 12409: 0x1200, 206: 0x13c0, 23669: 0x2302, 44: 0x120, 932: 0x762, 43: 0xb62, 21319: 0x1522, 9552: 0x1b82, 19974: 0x2042, 20059: 0x2742, 20063: 0x2ac0, 20117: 0x40, 58710: 0x162, 26518: 0x1160, 197: 0x1bc0, 26080: 0x1f42, 931: 0xcc0, 12421: 0x10a2, 111: 0x2820, 52: 0x2720, 954: 0x2a62, 51: 0x6c0, 21491: 0xae0, 9511: 0x12e0, 9489: 0x21e2, 12359: 0x2620, 181: 0x7c0, 955: 0x1260, 105: 0x20c0, 9580: 0x2660, 58706: 0x2882, 12364: 0x18a0, 9595: 0x1ba0, 53: 0x2022, 12418: 0x22, 88: 0x7a0, 171: 0x980, 232: 0xc22, 956: 0x10a0, 234: 0x1820, 224: 0x1d82, 58700: 0x2062, 12416: 0x7e0, 231: 0xd20, 12387: 0xfa0, 126: 0x10c0, 96: 0x1780, 167: 0x782, 216: 0x1360, 9596: 0x1fa2, 9522: 0xba0, 9628: 0xc62, 239: 0x1b80, 21502: 0x1f02, 12413: 0x2940, 31354: 0x6a0, 9588: 0xb22, 9512: 0x22a2, 19981: 0x22c2, 198: 0x29e0, 9567: 0x1840, 20058: 0x1a60, 9630: 0x1da2, 169: 0x2362, 255: 0x2540, 20061: 0x4c2, 124: 0x882, 253: 0x12a2, 61: 0x2a20, 99: 0x1782, 58707: 0x1a00, 12377: 0x23c2, 9504: 0x542, 9492: 0x5a0, 67: 0x9c2, 12360: 0xa00, 9495: 0xda0, 65: 0x26a2, 75: 0x28a0, 221: 0x2240, 46: 0x2c0, 12424: 0xfc2, 29976: 0x10e0, 21322: 0x1242, 12430: 0x1ea2, 22828: 0x42, 24187: 0xc00, 69: 0x1482, 9480: 0x2260, 12370: 0x23c0, 186: 0x23e2, 9599: 0x2662, 12433: 0xb40, 209: 0xc02, 58712: 0x1540, 26085: 0x1b62, 199: 0x2380, 30006: 0x1b22, 93: 0x1de0, 233: 0x662, 967: 0xcc2, 21489: 0xd42, 9583: 0x1902, 59: 0x1942, 205: 0x29e2, 58713: 0x2a42, 9574: 0x1a0, 926: 0x502, 21323: 0x640, 120: 0xb00, 12434: 0x2040, 9525: 0x2862, 24182: 0x2b02, 9590: 0xa80, 201: 0x18e2, 12356: 0x23a2, 959: 0x2762, 251: 0x2860, 195: 0x18e0, 86: 0x1a02, 58692: 0x28e2, 20049: 0x160, 19983: 0x7e2, 248: 0xc82, 229: 0x1420, 21318: 0x16e0, 22234: 0x2b00, 9474: 0xe0, 23665: 0x1042, 9518: 0x1620, 915: 0x19c2, 200: 0x2142, 21488: 0x2582, 9499: 0x0, 161: 0x1442, 965: 0x1cc0, 30335: 0x2122, 964: 0x2282, 242: 0x2982, 20011: 0x220, 20005: 0x1240, 12422: 0x12e2, 9565: 0x1b60, 9556: 0x2200, 58691: 0x1c0, 9500: 0x7c2, 20115: 0xce2, 204: 0x1180, 946: 0x27a2, 210: 0x660, 30333: 0x780, 9535: 0xa22, 165: 0xca0, 174: 0x7a2, 26086: 0x1320, 9531: 0x1e02, 26519: 0x1e80, 9475: 0x27e2, 21500: 0x240, 9559: 0x1822, 20113: 0x2120, 20048: 0x2400, 58703: 0x2a40, 9601: 0x1e82, 237: 0x9c0, 12404: 0xf40, 9486: 0x11c2, 20014: 0x17c0, 240: 0x1be0, 58708: 0x2222, 20101: 0x27a0, 92: 0x540, 9534: 0x5e0, 9584: 0xa62, 39: 0x1700, 189: 0x1bc2, 21496: 0x1dc0, 50: 0x29a0, 12414: 0x6e0, 966: 0xba2, 9491: 0xe82, 12365: 0xf82, 9586: 0x1c60, 9572: 0x302, 48: 0x1c40, 110: 0x2280, 12392: 0x2462, 30334: 0x2482, 82: 0x2622, 921: 0x2640, 26409: 0x102, 913: 0x1c2, 9564: 0x682, 9482: 0x1222, 9582: 0x15c2, 58709: 0x2900, 244: 0x2ae0, 19980: 0x4c0, 24181: 0xee2, 21321: 0x1302, 178: 0x13e2, 72: 0x1762, 9602: 0x2100, 184: 0x26c0, 109: 0x2840, 58701: 0x2a2, 21324: 0x342, 9476: 0x702, 12353: 0xa40, 9540: 0x10c2, 9620: 0x3a0, 922: 0x5c0, 9554: 0xc60, 24038: 0x1d62, 176: 0x29c2, 929: 0x25c2, 20052: 0x28c2, 34: 0x2b20, 12412: 0x20, 918: 0x182, 9561: 0x400, 9506: 0x1462, 24186: 0x1c22, 12362: 0x560, 30002: 0x1440, 9548: 0x19a0, 250: 0x27c0, 115: 0xc42, 26525: 0xde0, 402: 0x1060, 9477: 0x1882, 76: 0x1520, 9609: 0x1722, 218: 0x62, 9598: 0x9e2, 12406: 0x1c20, 170: 0x2962, 9536: 0x2aa0, 21503: 0x860, 9528: 0xbc0, 33: 0x1860, 21315: 0x1ac2, 9549: 0x802, 9553: 0xe22, 9625: 0x1920, 933: 0x19a2, 33457: 0x26a0, 20118: 0x15e2, 87: 0x1f60, 241: 0x21a2, 9516: 0x2a60, 73: 0x1ca2, 12372: 0x622, 123: 0xf80, 219: 0x1402, 78: 0x1ae0, 117: 0x1c42, 12393: 0x11e2, 12379: 0x1640, 21320: 0x1802, 166: 0x2e2, 9505: 0x382, 963: 0x4a0, 191: 0xca2, 12401: 0xec2, 193: 0x1862, 24040: 0x27c2, 20000: 0x940, 20111: 0xb42, 9627: 0x1422, 71: 0x17e0, 12391: 0x2682, 58693: 0x22e0, 230: 0x2980, 22827: 0xe62, 928: 0x1082, 9496: 0x11e0, 20109: 0x16a0, 162: 0x1ee0, 12432: 0x922, 254: 0xe02, 22232: 0x2220, 98: 0xaa2, 212: 0xc20, 114: 0xf00, 20110: 0x2202, 12354: 0x822, 208: 0x1f62, 30003: 0x4e2, 227: 0x12c0, 214: 0x8e2, 20050: 0x17c2, 9502: 0x22a0, 26410: 0x2060, 24178: 0x2180, 20060: 0x2422, 89: 0x602, 58704: 0x6a2, 20012: 0x6e2, 9613: 0x1102, 19968: 0x1ce2}
var eightData = []uint32{0x18, 0xff0018, 0xff0018, 0xff001f, 0xff001f, 0xff0000, 0xff0000, 0xff0000, 0x2009d, 0xf0051, 0x2001d, 0xf0011, 0x620039, 0x420055, 0x3c0009, 0x0, 0x80014, 0x3e0014, 0x28007f, 0x7f0014, 0x8007f, 0x140014, 0x630012, 0x0, 0x700007, 0x0, 0x66001c, 0x660036, 0x660063, 0x66007f, 0x3c0063, 0x0, 0x3f0000, 0x330000, 0x300000, 0x180011, 0xc0021, 0xc0025, 0xc0002, 0x0, 0x8, 0x220008, 0x790008, 0x2100f8, 0x2100ff, 0x220018, 0x100018, 0x18, 0x440032, 0x440002, 0x220027, 0x110022, 0x220072, 0x440029, 0x440011, 0x0, 0x700008, 0x8, 0x180008, 0x3c0008, 0x660008, 0x3c0008, 0x180008, 0x8, 0x80008, 0x3e002a, 0x8003e, 0x1c0000, 0x2a007f, 0x490048, 0xc0064, 0x0, 0x0, 0x0, 0x6e0000, 0x3b0000, 0x130000, 0x3b000c, 0x6e000c, 0x6, 0x140000, 0x140000, 0x140000, 0x140000, 0x17000f, 0x140008, 0x140008, 0x140008, 0x7f, 0x8, 0x18007a, 0x180019, 0x7e0028, 0x2a, 0x7e0049, 0x0, 0x7f001e, 0x630006, 0x310006, 0x180006, 0x4c0006, 0x660006, 0x7f001e, 0x0, 0x80000, 0x80000, 0x7f0000, 0x800ff, 0x140000, 0x2200f7, 0x490014, 0x14, 0xc0000, 0x1e0000, 0x33000c, 0x330000, 0x3f000c, 0x330030, 0x330030, 0x0, 0x6e00c3, 0x3b0063, 0x33, 0x1e00bd, 0x3300ec, 0x3300f6, 0x1e00f3, 0x3, 0x7f0000, 0x80000, 0x18001f, 0x280033, 0x80033, 0x80033, 0x80033, 0x30, 0x3e0041, 0x300022, 0x140014, 0x7f0008, 0x80008, 0x80008, 0x80008, 0x0, 0x60077, 0xc0045, 0x180047, 0x180050, 0x180048, 0xc0044, 0x60060, 0x0, 0x7e0018, 0xc, 0x6, 0x3, 0x6, 0xc, 0x18, 0x0, 0x4002d, 0x7e0000, 0x10033, 0x3c0033, 0x100033, 0x80033, 0x7c001e, 0x0, 0x7f, 0x2, 0x63003e, 0x600022, 0x60003e, 0x600020, 0x7f007f, 0x0, 0x40000, 0x1f0000, 0x150000, 0x1f0000, 0x150000, 0x5f000c, 0x7c000c, 0x0, 0x180000, 0x180000, 0x180000, 0xf, 0x18000f, 0x180000, 0x180000, 0x0, 0x3, 0xc4, 0x63, 0xff00b4, 0xdb, 0xff00ac, 0x800e6, 0x80080, 0x200000, 0x240000, 0x24006e, 0x340033, 0x2c0033, 0x24003e, 0x200030, 0x78, 0x55001c, 0x550010, 0x550014, 0x7f0014, 0x550022, 0x550022, 0x550041, 0x0, 0x420000, 0x2f0000, 0x20000, 0x720018, 0x20000, 0x90066, 0x710000, 0x0, 0x180000, 0x180000, 0x18000c, 0xf8000c, 0xf8000c, 0x8002c, 0x80018, 0x80000, 0x2200ff, 0x7f0000, 0x140000, 0x550000, 0x550000, 0x140000, 0x7f0000, 0x0, 0x0, 0x3f0000, 0x6007e, 0x30000, 0x3007e, 0x1e0018, 0x300018, 0x1c0000, 0x1e0033, 0x30033, 0xe0033, 0x30033, 0x30033, 0x1e0033, 0x30003f, 0x1c0000, 0x80014, 0x7f0014, 0x80014, 0x3e0014, 0x1400fc, 0x2a0000, 0x490000, 0x0, 0x33003f, 0x2d, 0x1e000c, 0xc000c, 0xc000c, 0xc000c, 0x1e001e, 0x0, 0x80, 0x80, 0x130080, 0x320080, 0x510080, 0x110080, 0xe0080, 0x80, 0x1f0000, 0x140000, 0x140008, 0x1f0008, 0x240010, 0x240030, 0x44000c, 0x0, 0x6e0000, 0x3b0000, 0x3f, 0x3e0019, 0x63000c, 0x630026, 0x3e003f, 0x0, 0x80000, 0x90000, 0x3e007e, 0x4b001b, 0x65001b, 0x55001b, 0x22000e, 0x0, 0x4007f, 0x40014, 0x1f0014, 0x140014, 0x120014, 0x520012, 0x710011, 0x0, 0x8007f, 0x3e0008, 0x2a007f, 0x3e0008, 0x2a0008, 0x3e0014, 0x80063, 0x0, 0x7f0000, 0x63000c, 0xc, 0x3e0000, 0x0, 0x63000c, 0x7f000c, 0x0, 0x240008, 0x240014, 0x2f0022, 0x240049, 0x2e0008, 0x350008, 0x640008, 0x0, 0x180003, 0x180006, 0x18000c, 0x180018, 0xf80030, 0x180060, 0x180040, 0x180000, 0x24, 0x4f, 0x40004, 0xa003c, 0x110046, 0x200045, 0x400022, 0x0, 0x80068, 0xa0008, 0x7f007f, 0x9001c, 0x3e002a, 0x80049, 0x7f0008, 0x0, 0xfe0008, 0xdb0008, 0xdb0008, 0xde0008, 0xd800f8, 0xd80000, 0xd80000, 0x0, 0x80067, 0x80066, 0x80036, 0x8001e, 0xf0036, 0x80066, 0x80067, 0x80000, 0x1e0008, 0x8, 0x20008, 0x3a00f8, 0x4600ff, 0x420008, 0x300008, 0x8, 0x33007f, 0x330041, 0x33005e, 0x1e0052, 0xc005e, 0xc0052, 0x1e005e, 0x0, 0x200010, 0x400008, 0x160004, 0x200002, 0x10004, 0x10008, 0xe0010, 0x0, 0x80008, 0x7f007f, 0x410008, 0x1c0022, 0x10007f, 0x140022, 0x22003e, 0x0, 0x38000e, 0x0, 0x1e0018, 0x33003c, 0x3f0066, 0x3003c, 0x1e0018, 0x0, 0x140008, 0x140008, 0x140008, 0x1400f8, 0x1f00ff, 0x0, 0x0, 0x0, 0x8, 0x7f, 0x660055, 0x600000, 0x66003e, 0x60008, 0x66007f, 0x0, 0x1e, 0x220033, 0x7a0030, 0x22001c, 0x720030, 0x2a0033, 0x12001e, 0x0, 0x10007e, 0x120008, 0x14003e, 0x100008, 0x18001c, 0x14002a, 0x120004, 0x0, 0x3e, 0x22, 0x3e, 0x22, 0xbb003e, 0x22, 0x21, 0x0, 0x0, 0x0, 0x0, 0xff0000, 0xff00ff, 0x0, 0x0, 0x0, 0x180000, 0x180000, 0x12, 0x7e001f, 0x22, 0x180012, 0x180004, 0x0, 0x3f0033, 0x2d0000, 0xc001e, 0xc0033, 0xc003f, 0xc0003, 0x1e001e, 0x0, 0x7c0004, 0xc60008, 0x1c003e, 0x360022, 0x36003e, 0x1c0022, 0x33003e, 0x1e0000, 0x3c0063, 0x420063, 0x9d0036, 0xa5001c, 0x9d001c, 0xa50036, 0x420063, 0x3c0000, 0x80000, 0x80000, 0x80066, 0x80066, 0xf80066, 0x8003e, 0x80006, 0x80003, 0x7f0004, 0x8000f, 0x3a0064, 0x2a0006, 0x3e0005, 0x200026, 0x30003c, 0x0, 0x3e, 0x63, 0x7b, 0xe7007b, 0xe7007b, 0x3, 0x1e, 0x0, 0x40018, 0x3f0018, 0x40018, 0x3c0000, 0x560000, 0x4d0018, 0x260018, 0x18, 0x80080, 0x7f0050, 0x49003a, 0x490017, 0x7f001a, 0x80002, 0x8001c, 0x0, 0x3e0000, 0x40077, 0x50025, 0x7f0025, 0x140027, 0x520020, 0x710070, 0x0, 0x180008, 0x180008, 0x180008, 0xff, 0x1800ff, 0x180018, 0x180018, 0x18, 0x8, 0x1c, 0x1c, 0x1f0036, 0xff0036, 0x180063, 0x18007f, 0x180000, 0x7e0068, 0x1c, 0x6b, 0x3c001c, 0x6b, 0x1c, 0x7f000b, 0x0, 0xc30036, 0x180036, 0x3c007f, 0x660036, 0x66007f, 0x3c0036, 0x180036, 0x0, 0x0, 0x18, 0x0, 0x18, 0xff0000, 0x140018, 0x140000, 0x140018, 0xe0008, 0x80008, 0x3c0008, 0x4a0008, 0x690018, 0x550018, 0x320018, 0x18, 0x1e0014, 0x330014, 0x300077, 0x180014, 0xc0077, 0x0, 0xc003e, 0x0, 0x2007f, 0xe0008, 0x80008, 0x7f0008, 0x80008, 0x140008, 0x630008, 0x0, 0x180000, 0x1800cc, 0x180066, 0xf80033, 0xff0066, 0xcc, 0x0, 0x0, 0x80008, 0x140008, 0x22003e, 0x490008, 0x80008, 0x490008, 0x7f007f, 0x0, 0x3c001c, 0x660000, 0x3000e, 0x3000c, 0x3000c, 0x66000c, 0x3c001e, 0x0, 0x18, 0x18, 0x18, 0xf0018, 0xff001f, 0x8, 0x8, 0x8, 0x7f001c, 0x140000, 0x7f003e, 0x550010, 0x7f0038, 0x140024, 0x7f0062, 0x0, 0x80006, 0x8000c, 0x80018, 0xff0030, 0xff0018, 0x8000c, 0x80006, 0x80000, 0x7e0004, 0xc3003f, 0x3c0004, 0x60003c, 0x7c0056, 0x66004d, 0xfc0026, 0x0, 0x80000, 0x80000, 0x80000, 0x1000f8, 0xe000f8, 0x18, 0x18, 0x18, 0xf00000, 0xf00000, 0xf00000, 0xf00000, 0xff00f8, 0xff0000, 0xff0000, 0xff0000, 0x7003f, 0x60010, 0x60008, 0x3e003c, 0x660042, 0x660041, 0x3b0030, 0x0, 0x7f0038, 0x2000c, 0x20018, 0x3f003e, 0x220033, 0x220033, 0x7f001e, 0x0, 0x7f0004, 0x63007f, 0x310004, 0x18007a, 0x4c0049, 0x660048, 0x7f0078, 0x0, 0xf0000, 0x60000, 0x3e0063, 0x660036, 0x66001c, 0x3e0036, 0x60063, 0xf0000, 0x33, 0x0, 0x1e, 0x30, 0xf003e, 0x33, 0x7e, 0x0, 0x3e0006, 0x3c, 0x7f0042, 0x40039, 0x3c0004, 0x200036, 0x300049, 0x0, 0x0, 0xc0000, 0xc0018, 0x3f0000, 0xc0018, 0xc0000, 0x18, 0x0, 0x1c0080, 0x51, 0x8001d, 0x2a0011, 0x490039, 0x100015, 0xc0009, 0x0, 0x0, 0x0, 0x760000, 0xdb00f8, 0xdb00ff, 0x7e0018, 0x180018, 0x18, 0x140018, 0x140018, 0x140018, 0xf40018, 0x400ff, 0xf40000, 0x140000, 0x140000, 0x2004f, 0xf0024, 0x20002, 0x72007f, 0x20008, 0x90004, 0x710038, 0x0, 0x3f0074, 0x42, 0x330054, 0x370048, 0x3f0044, 0x3b0052, 0x33005f, 0x0, 0x7003c, 0x66, 0x1e0018, 0x33003c, 0x3f0066, 0x3003c, 0x1e0018, 0x0, 0x0, 0x0, 0x3e0000, 0x30000, 0x1e00ff, 0x300018, 0x1f0018, 0x18, 0xff0000, 0xff0000, 0xff0000, 0xff00f8, 0xf00008, 0xf000f8, 0xf00008, 0xf00008, 0x3c, 0x600000, 0x3c003c, 0x760042, 0x7e0040, 0x6e0020, 0x3c0018, 0x60000, 0xc0033, 0x33, 0xc001e, 0x6003f, 0x3000c, 0x33003f, 0x1e000c, 0xc, 0x7f, 0x630063, 0x360006, 0x1c000c, 0x1c0006, 0x360063, 0x63007f, 0x0, 0x3e0014, 0x14, 0x14, 0x7f0014, 0x2200ff, 0x220014, 0x210014, 0x14, 0x7007f, 0x60040, 0x36005f, 0x6e0055, 0x66005f, 0x660055, 0x67005f, 0x0, 0x380000, 0x0, 0xc001e, 0xc0003, 0xc0003, 0x2c001e, 0x180030, 0x1c, 0x100044, 0x10002f, 0x770004, 0x15001e, 0x150020, 0x570020, 0x70001c, 0x0, 0x3e007f, 0x80041, 0x6b0041, 0x80041, 0x7f0041, 0x80041, 0x8007f, 0x0, 0x700000, 0x0, 0x6e0000, 0x3b003f, 0x130000, 0x3b0000, 0x6e0000, 0x0, 0x800018, 0x400018, 0x200018, 0x1000f8, 0x800f8, 0x40000, 0x20000, 0x10000, 0x22007f, 0x2a0002, 0x2a007f, 0x7f0002, 0x2a003e, 0x2a0020, 0x210038, 0x0, 0x22, 0x72, 0x380027, 0x660072, 0x60057, 0x6002a, 0x70052, 0x0, 0xc, 0x12, 0x6003f, 0x3e0006, 0x66001e, 0x3e0006, 0x6003f, 0x0, 0x140033, 0x140033, 0x140033, 0x14003f, 0x140033, 0x140033, 0x140033, 0x140000, 0x40008, 0x240008, 0x4f0008, 0x54000f, 0x520008, 0x12000f, 0x90008, 0x8, 0x80014, 0x7f0022, 0x8007f, 0x7f0002, 0x8003e, 0x140020, 0x630030, 0x0, 0x63, 0x77, 0x7f, 0x1f007f, 0x1f006b, 0x180063, 0x180063, 0x180000, 0xf0000, 0xf0000, 0xf0000, 0xf00ff, 0xff, 0xff, 0xff, 0xff, 0x400040, 0xb10020, 0x5d001e, 0x110021, 0x390020, 0x150020, 0x9001c, 0x0, 0x77000e, 0x220008, 0x220008, 0x7f0014, 0x220014, 0x220022, 0x220041, 0x0, 0x0, 0x330000, 0x3b, 0x33006e, 0x330066, 0x330006, 0x7e000f, 0x0, 0xff0018, 0xff0018, 0xff0018, 0xff0018, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0x80040, 0x800a0, 0x80043, 0x7f0032, 0x80051, 0x80011, 0x8000e, 0x0, 0x1c0014, 0x6007f, 0x30008, 0x1f002a, 0x33003e, 0x330008, 0x1e0006, 0x0, 0x80038, 0x1f000c, 0x8000c, 0x3f0007, 0x1c000c, 0x2000c, 0x3c0038, 0x0, 0x200000, 0x500000, 0x240000, 0xa001c, 0x110022, 0x200020, 0x40001c, 0x0, 0x80000, 0x80000, 0x380063, 0x8006b, 0x1c007f, 0x2a007f, 0x40036, 0x0, 0x4007f, 0x7f0055, 0x610055, 0x550055, 0x490073, 0x570041, 0x7f007f, 0x0, 0x60, 0x30, 0x330018, 0x33000c, 0x1e0006, 0xc0003, 0xc0001, 0x0, 0x30000, 0x30000, 0x30000, 0x30000, 0x30008, 0x30008, 0x30008, 0x30008, 0x80077, 0x80045, 0x490055, 0x490027, 0x490030, 0x490048, 0x7f0044, 0x0, 0xc0070, 0xe00d8, 0xc0018, 0xc003c, 0xc0018, 0xc0018, 0x3f001b, 0xe, 0x7f0018, 0x360018, 0x36007e, 0x360018, 0x360018, 0x360000, 0x36007e, 0x0, 0x0, 0x0, 0x110066, 0x3d0066, 0x530066, 0x39003e, 0x110006, 0x3, 0x18006e, 0x18003b, 0x180000, 0xf80000, 0xff0000, 0x80000, 0x80000, 0x80000, 0x180014, 0x180014, 0x18007f, 0x180014, 0xff001c, 0x180014, 0x18001c, 0x180000, 0x70008, 0x70008, 0x70008, 0x70008, 0x700ff, 0x70018, 0x70018, 0x70018, 0x0, 0x90000, 0x3d0000, 0x4a0000, 0x4b00f0, 0x4500f0, 0x2a00f0, 0xf0, 0x7f003e, 0x410002, 0x5d003e, 0x590022, 0x7f007f, 0x490000, 0x7f0030, 0x0, 0x1e0042, 0x330037, 0x330022, 0x3e0077, 0x30002a, 0x180052, 0xe0052, 0x0, 0x80007, 0x80000, 0x7f001e, 0xc, 0x14000c, 0x22000c, 0x41001e, 0x0, 0x630000, 0x770000, 0x7f006e, 0x7f0033, 0x6b0033, 0x63003e, 0x630030, 0x1f, 0x0, 0x0, 0x7f, 0x36, 0xf80036, 0x180036, 0x180036, 0x180000, 0x440008, 0x240008, 0x80008, 0x3c0008, 0x2000f, 0x20000, 0x3c0000, 0x0, 0x80020, 0x7f0040, 0x80014, 0xa002a, 0xc0011, 0xc0020, 0x330040, 0x0, 0x3f, 0x80066, 0x66, 0x8003e, 0x66, 0x80066, 0x3f, 0x80000, 0x6b007f, 0x80014, 0x3e0055, 0x80014, 0x7f007f, 0x80001, 0x80001, 0x0, 0x0, 0x3, 0x1e0006, 0x3000c, 0xe001c, 0x30036, 0x1e0063, 0x0, 0x80000, 0x7f0000, 0x49001e, 0x7f0033, 0x18003f, 0x240033, 0x42001e, 0x0, 0x7f, 0x380001, 0x1d, 0x330015, 0x33001d, 0x3e0001, 0x30007f, 0x1f0000, 0x38006e, 0x3b, 0x1e001e, 0x30030, 0xe003e, 0x30033, 0x1e007e, 0x0, 0x8, 0x110008, 0x3d0008, 0x530008, 0x51001f, 0x390018, 0x110018, 0x18, 0x80008, 0x3e0008, 0x80008, 0x220038, 0x7f0008, 0x220008, 0x21007f, 0x0, 0x7003e, 0xc0022, 0xc003e, 0x380022, 0xc003e, 0xc0000, 0x7007f, 0x0, 0x40, 0x2e, 0x300000, 0x30003c, 0x3e0042, 0x300040, 0x300038, 0x0, 0xdb005c, 0xdb0036, 0xdb0073, 0x7e007b, 0x18006f, 0x180036, 0x3c001d, 0x0, 0x4, 0x7, 0xc0034, 0xc004c, 0x180066, 0x300054, 0x7f0024, 0x0, 0x0, 0x0, 0x180066, 0x3c0060, 0x660030, 0x660018, 0x66000c, 0x0, 0x7f000c, 0x460012, 0x160000, 0x1e001e, 0x16000c, 0x46000c, 0x7f001e, 0x0, 0x1c00ff, 0x3000aa, 0x1800ff, 0xc0055, 0x3c00ff, 0xaa, 0xff, 0x55, 0x3c0000, 0x660036, 0x1c, 0x660008, 0x66001c, 0x660036, 0x3c0000, 0x0, 0xff000c, 0xff000c, 0xff001e, 0xff0030, 0xf003e, 0xf0033, 0xf007e, 0xf0000, 0x18003e, 0x18002a, 0x3e, 0x18002a, 0x18003e, 0x180008, 0x180008, 0x0, 0x80000, 0x80000, 0x80063, 0xf8003e, 0xf80036, 0x18003e, 0x180063, 0x180000, 0x7f0000, 0x460000, 0x160018, 0x1e0018, 0x160000, 0x460018, 0x7f0018, 0x0, 0x7f, 0x46, 0x330016, 0x33001e, 0x330016, 0x1e0006, 0xc000f, 0x0, 0x18, 0x0, 0x18, 0x18, 0xfc0018, 0x140000, 0x140018, 0x140018, 0x400024, 0x20004f, 0x30014, 0x32002e, 0x510001, 0x110001, 0xe000e, 0x0, 0x30001c, 0x270036, 0x25001c, 0x27006e, 0x30003b, 0x480033, 0x44006e, 0x0, 0x14000f, 0x130006, 0x120006, 0x7f0006, 0x120046, 0x120066, 0x12007f, 0x0, 0x7f0000, 0x140000, 0x770066, 0x410030, 0x770018, 0x14000c, 0x7f0006, 0x0, 0x380008, 0x30001c, 0x30001c, 0x3e0036, 0x330036, 0x330063, 0x6e0063, 0x0, 0x38001c, 0x36, 0x1e0000, 0x30003e, 0x3e0063, 0x33007f, 0x7e0063, 0x0, 0x10004c, 0x1400a0, 0x140048, 0x7f000a, 0x140029, 0x140048, 0x10000c, 0x0, 0x1c, 0x36, 0x6, 0xf, 0x30006, 0x40006, 0x8000f, 0x80000, 0x3e0004, 0x4, 0x7f007f, 0x4, 0x3e0004, 0x24, 0x7f003c, 0x0, 0x140008, 0x140008, 0x14003e, 0x17002a, 0x10003e, 0x170022, 0x14003e, 0x140000, 0x0, 0x0, 0x0, 0xf8, 0xe000ff, 0x100008, 0x80008, 0x80008, 0x200022, 0x2e007f, 0x2a0022, 0x7a0022, 0x2a001a, 0x2e0002, 0x20001c, 0x0, 0x70008, 0x7f, 0xe0055, 0xc0000, 0xc001e, 0xc0008, 0x1e003c, 0x0, 0x180036, 0x180000, 0x18003f, 0xf80006, 0xf8001e, 0x180006, 0x18003f, 0x180000, 0x3e, 0x630000, 0x33007f, 0x180008, 0xc0008, 0x660008, 0x630008, 0x0, 0x200000, 0x400000, 0x140036, 0x240063, 0x8006b, 0x18007f, 0x60036, 0x0, 0x80002, 0x8001f, 0x7f0052, 0x80061, 0x80008, 0x14007f, 0x630008, 0x0, 0x80006, 0x80006, 0x80003, 0xf80000, 0xf80000, 0x0, 0x0, 0x0, 0x7f0000, 0x7f0000, 0x7f0000, 0x7f0000, 0x7f00ff, 0x7f00ff, 0x7f00ff, 0x7f00ff, 0x7f007f, 0x410045, 0x5d0055, 0x55005d, 0x5d0051, 0x410041, 0x7f007f, 0x0, 0x33003e, 0x330063, 0x33001c, 0x3f0018, 0x330018, 0x330018, 0x33003c, 0x0, 0xc, 0xc, 0x1e0018, 0x330000, 0x30000, 0x330000, 0x1e0000, 0x0, 0x400000, 0xa00000, 0x400000, 0x0, 0xf8, 0x8, 0x8, 0x8, 0x3e001f, 0x20012, 0x3e0017, 0x220012, 0x7f0017, 0x52, 0x60062, 0x0, 0x2003c, 0x3f0066, 0x80003, 0xe0003, 0xa0073, 0x7f0066, 0x8007c, 0x0, 0x20072, 0x3e0057, 0x90076, 0x80056, 0x7f0075, 0x80005, 0x80079, 0x0, 0x7e, 0xc3, 0x3c, 0x1f0066, 0x10007e, 0x170006, 0x14003c, 0x140000, 0x14, 0x80014, 0x380014, 0x80014, 0x1c00f4, 0x2a0014, 0x40014, 0x14, 0x700018, 0x3c, 0x1c003c, 0x360018, 0x630018, 0x7f0000, 0x630018, 0x0, 0x33, 0x0, 0x33, 0xbb0033, 0xbb0033, 0x33, 0x1e, 0x0, 0x44, 0x20024, 0x7a000f, 0x20054, 0xa0052, 0x720052, 0x20009, 0x0, 0x8, 0x8, 0x7f, 0x1c, 0x2a, 0x49, 0x8, 0xff0000, 0x38006e, 0x3b, 0x3f0000, 0x6003e, 0x1e0063, 0x6007f, 0x3f0063, 0x0, 0x80008, 0x80008, 0x80008, 0x400f8, 0x300f8, 0x8, 0x8, 0x8, 0xf, 0x7000f, 0x24000f, 0x24000f, 0x7e00ff, 0x2500ff, 0x1200ff, 0xff, 0x8, 0xc0008, 0xc0004, 0xc, 0x56, 0xc0052, 0xc0021, 0x60000, 0x80008, 0x2a003e, 0x3e002a, 0x8007f, 0x7f0008, 0x480014, 0x640063, 0x0, 0x54, 0x54, 0x2a, 0x3f0015, 0x3002a, 0x30054, 0x54, 0x0, 0x330000, 0x0, 0x330000, 0x330000, 0x1e00e7, 0xc0000, 0x1e0000, 0x0, 0x3f0008, 0x330008, 0x30008, 0x3000f, 0x300ff, 0x30008, 0x30008, 0x8, 0x7f0000, 0x20000, 0x1e0000, 0x1000f8, 0x1000f8, 0x100008, 0x180008, 0x8, 0x330000, 0x330000, 0x33000c, 0x33000c, 0x33003c, 0x1e0030, 0xc0030, 0x0, 0x180018, 0x180018, 0x180018, 0xf800ff, 0xff00ff, 0x180018, 0x180018, 0x180018, 0x550014, 0xaa0014, 0x550014, 0xaa0014, 0x5500ff, 0xaa0000, 0x550000, 0xaa0000, 0x1, 0x3e0001, 0x80001, 0x80001, 0x80001, 0x80041, 0x7f007f, 0x0, 0x10001e, 0x7e0033, 0x100033, 0x3c001e, 0x20033, 0x20033, 0x1c001e, 0x0, 0x1d0000, 0x110012, 0x11003f, 0x290042, 0x450042, 0x10034, 0x7f0004, 0x0, 0x600022, 0x1c0014, 0xb0000, 0x8007f, 0x7f0048, 0x80068, 0x80008, 0x0, 0xff0063, 0xff0067, 0xff006f, 0xff007b, 0x73, 0x63, 0x63, 0x0, 0x7f007e, 0x400000, 0x200000, 0x10003c, 0x100000, 0x100000, 0x18007f, 0x0, 0x300079, 0xc0049, 0x3e0049, 0x2a0079, 0x3e0049, 0x2a0049, 0x3e0079, 0x0, 0x0, 0x66, 0x3c, 0xff, 0x1f003c, 0x140066, 0x140000, 0x140000, 0x3e0014, 0x220014, 0x220014, 0x3e0017, 0x220010, 0x22001f, 0x3e0000, 0x0, 0x33, 0x0, 0xe, 0xff000c, 0xc, 0xff000c, 0x1e, 0x0, 0x0, 0x20000, 0x20000, 0x20000, 0x420018, 0x220018, 0x1c0018, 0x18, 0xc3000c, 0x63000c, 0x330000, 0x7b001e, 0xcc0033, 0x66003f, 0x330033, 0xf00000, 0x1b, 0x33000e, 0x1b, 0x1e0030, 0x33003e, 0x330033, 0x1e001e, 0x0, 0x0, 0x3e0000, 0x80000, 0x40000, 0x4001f, 0x40018, 0x380018, 0x18, 0x8004c, 0x450020, 0x220008, 0x14002a, 0x80049, 0x440010, 0x7f000c, 0x0, 0x3e, 0x63, 0x330073, 0x33007b, 0x33006f, 0x330067, 0x6e003e, 0x0, 0x1, 0x2, 0x4, 0x8, 0xf0010, 0xf0020, 0xf0040, 0xf0080, 0x1f0000, 0x80000, 0x3c003c, 0x420000, 0x490002, 0x540002, 0x38003c, 0x0, 0x1e001f, 0xc0033, 0xc0033, 0xc005f, 0xc0063, 0xc00f3, 0x1e0063, 0xe3, 0x0, 0x0, 0x33, 0xff0033, 0xff0033, 0x80033, 0x8001e, 0x80000, 0x0, 0x0, 0x0, 0x7f000f, 0x8, 0xf, 0x8, 0x8, 0x8007f, 0x80014, 0x2a0077, 0x490055, 0x80077, 0x80014, 0xc0012, 0x0, 0x300000, 0x0, 0x300000, 0x3000ff, 0x300000, 0x3300db, 0x330000, 0x1e0000, 0x0, 0x0, 0x3c003b, 0x660066, 0x660066, 0x36003e, 0x60006, 0x6000f, 0x40008, 0x7f0008, 0x40008, 0x3c0008, 0x140008, 0x120008, 0x7f007f, 0x0, 0x7001c, 0x36, 0x1e0063, 0x300063, 0x3e0063, 0x330036, 0x7e001c, 0x0, 0xf00008, 0xf00008, 0xf00008, 0xf000ff, 0xf00ff, 0xf0000, 0xf0000, 0xf0000, 0x68007f, 0x80040, 0x7f004e, 0x80040, 0x2a004e, 0x2a004a, 0x69006e, 0x0, 0x8001e, 0x80018, 0x80018, 0x1f0018, 0xff0018, 0x180018, 0x18001e, 0x180000, 0x180028, 0x180044, 0x180012, 0xff0021, 0xff0002, 0x4, 0x8, 0x0, 0x0, 0x33, 0x3c0066, 0x4a00cc, 0x490066, 0x450033, 0x220000, 0x0, 0x80078, 0x7e0010, 0x20077, 0x20042, 0x20042, 0x20057, 0x10070, 0x0, 0x1c0077, 0x360055, 0x260077, 0xf0055, 0x60077, 0x670055, 0x3f004d, 0x0, 0x22, 0x22, 0x77, 0x22, 0x77, 0x2a, 0x22, 0xff0000, 0x18, 0x18, 0x80018, 0xe001f, 0x3800ff, 0x4c0000, 0x2a0000, 0x0, 0x80030, 0x8000f, 0x3e0008, 0x2a007f, 0x3e0008, 0x2a0014, 0x3e0063, 0x0, 0x1e0018, 0x80018, 0x4007e, 0x7f0003, 0x80003, 0x4007e, 0x380018, 0x18, 0x3e0008, 0x10007e, 0x80008, 0x8003c, 0x1c0040, 0x140040, 0x1c0038, 0x0, 0x3f, 0x66, 0x66, 0xf003e, 0xff0066, 0x80066, 0x8003f, 0x80000, 0x3e0008, 0x80008, 0x8007f, 0x7f0049, 0x180049, 0x540055, 0x730041, 0x0, 0x3f0063, 0x660063, 0x6f0063, 0x6f006b, 0x66007f, 0x660077, 0x3f0063, 0x0, 0xe, 0x380000, 0x66, 0x1e0066, 0x330066, 0x330066, 0x1e003c, 0x0, 0x1c, 0x36, 0x63, 0xf8007f, 0xff0063, 0x36, 0x1c, 0x0, 0x1c007f, 0x220008, 0x5d0008, 0x80008, 0x3e0008, 0x80008, 0x3e000c, 0x0, 0x1e0000, 0x330000, 0x7001e, 0xe0033, 0x38003f, 0x330003, 0x1e001e, 0x0, 0x18001c, 0x7e0036, 0xdb0063, 0xdb0063, 0xdb0063, 0x7e0036, 0x18001c, 0x0, 0x3f0008, 0x30010, 0x1f0022, 0x300049, 0x300024, 0x33003e, 0x1e0020, 0x0, 0x7f0004, 0x8000f, 0x40004, 0x2006e, 0x10011, 0x8, 0x70, 0x0, 0x8, 0x3e, 0x30008, 0x33007f, 0x30014, 0x33002a, 0x30049, 0x0, 0x0, 0x1c, 0x330000, 0x33001c, 0x330022, 0x3e0020, 0x300018, 0x1f0000, 0x7f0004, 0x41003c, 0x4f0004, 0x49001c, 0x4f0010, 0x41007f, 0x7f0010, 0x0, 0x63000c, 0x1c0000, 0x36000e, 0x63000c, 0x7f000c, 0x63000c, 0x63001e, 0x0, 0x810000, 0x420007, 0x240000, 0x180033, 0x180033, 0x240033, 0x42007e, 0x810000, 0xf0000, 0xf0000, 0xf0000, 0xf0000, 0xf00000, 0xf00000, 0xf000ff, 0xf000ff, 0x12003e, 0x170000, 0x15007f, 0x170002, 0x150022, 0x57003e, 0x700020, 0x0, 0x7000c, 0x3e, 0x3f0003, 0x6001e, 0x1e0030, 0x6001f, 0x3f000c, 0x0, 0x440000, 0x2f0000, 0x6003e, 0x50003, 0x60003, 0x4001e, 0x30030, 0x1c, 0x3e, 0x8, 0x180008, 0x18007f, 0x5a0008, 0x180008, 0x180008, 0x0, 0x0, 0x1f0000, 0x1e, 0x1f0030, 0x33003e, 0x330033, 0x33006e, 0x0, 0x3e, 0x63, 0xdb0063, 0xdb0063, 0xdb0036, 0x7e0036, 0x180077, 0x0, 0x1f, 0x36, 0x66, 0xf0066, 0xf0066, 0x80036, 0x8001f, 0x80000, 0x3e0000, 0x80000, 0x80000, 0x7f00fc, 0x80004, 0x800f4, 0x80014, 0x14, 0x7f, 0x41, 0x3c005d, 0x300055, 0x30005d, 0x45, 0x30005d, 0x0, 0x70, 0x1e0000, 0x330066, 0x1f0066, 0x33003c, 0x1f0018, 0x30018, 0x30000, 0x3e0000, 0x220000, 0x3e0000, 0x220000, 0x3e0055, 0x80000, 0x7f0000, 0x80000, 0x0, 0x0, 0x7e001f, 0x180033, 0x180033, 0x580033, 0x300033, 0x0, 0x180018, 0x180018, 0x180018, 0x180018, 0x1f00f8, 0x180008, 0x180008, 0x180008, 0x7f0008, 0x80008, 0x1c0008, 0x2a0008, 0x80000, 0x80000, 0x80000, 0x0, 0x330000, 0x0, 0x63, 0x3, 0x3, 0x3, 0x7f, 0x0, 0x8001f, 0x2a001f, 0x3e001f, 0x1f, 0x7e001f, 0x2001f, 0x1001f, 0x1f, 0x8002d, 0x80000, 0x8000c, 0x8000c, 0xff000c, 0x2c, 0x18, 0x0, 0x55, 0x1c0000, 0x2200aa, 0x410000, 0x400055, 0x200000, 0x1c00aa, 0x0, 0x3c0022, 0x42007f, 0x990022, 0x850008, 0x85003a, 0x99000a, 0x42007f, 0x3c0000, 0x8001e, 0x80033, 0x80003, 0xf80033, 0x8001e, 0xf80018, 0x30, 0x1e, 0x0, 0x10000, 0x110000, 0x210000, 0x210000, 0x2500ff, 0x200ff, 0xff, 0x100040, 0x7e0022, 0x180011, 0x14003d, 0x180011, 0x100012, 0xc0008, 0x0, 0x1c0008, 0x360008, 0x360008, 0x1c0008, 0x8, 0x8, 0x8, 0x0, 0x18003f, 0x180001, 0x180009, 0x1f007f, 0x1f0008, 0x18002a, 0x180049, 0x180000, 0x4003e, 0x3f0020, 0x240010, 0x340008, 0x40004, 0x440042, 0x7c007f, 0x0, 0x18, 0xc, 0x6, 0x180006, 0x180006, 0xc, 0x18, 0x0, 0x40008, 0x4000c, 0x80008, 0x3c001c, 0x20000, 0x20000, 0x3c0000, 0x0, 0x3e0044, 0x80022, 0x3e0077, 0x220022, 0x3e0077, 0x220022, 0x3e0022, 0x0, 0x63001e, 0x670033, 0x6f0033, 0x7b0033, 0x73003b, 0x63001e, 0x630038, 0x0, 0x200008, 0x420008, 0x120008, 0x22000f, 0x2000f, 0x220000, 0x1c0000, 0x0, 0x80008, 0x1c0008, 0x36007f, 0x63001c, 0x2a, 0x5d, 0x8, 0x0, 0x440008, 0x2f0008, 0x40008, 0x1f00ff, 0xe0000, 0x100ff, 0x1e0000, 0x0, 0x180004, 0x180007, 0x180004, 0x1f003c, 0xff0046, 0x80045, 0x80024, 0x80000, 0x80000, 0x80033, 0x80000, 0x80033, 0xff0033, 0x8003e, 0x80030, 0x8001f, 0x40014, 0x70014, 0x40014, 0xc00f7, 0x160000, 0x5500ff, 0x240000, 0x0, 0x40008, 0x22002a, 0x7f003e, 0x400000, 0x3e003e, 0x220000, 0x3e007f, 0x0, 0x78007f, 0x300049, 0x300049, 0x30007f, 0x330049, 0x330049, 0x1e007f, 0x0, 0x3f0000, 0x660000, 0x660033, 0x3e0033, 0x60033, 0x6001e, 0xf000c, 0x0, 0x80008, 0x80008, 0x80008, 0x1f0008, 0x1f00f8, 0x180018, 0x180018, 0x180018, 0xf00014, 0xf00014, 0xf00014, 0xf000f4, 0x4, 0xfc, 0x0, 0x0, 0x3f001c, 0x660000, 0x66003e, 0x3e0010, 0x360038, 0x660024, 0x670062, 0x0, 0x80033, 0x80000, 0x8001e, 0xff000c, 0x8000c, 0xff000c, 0x8001e, 0x80000, 0x180014, 0x180014, 0x180014, 0x1800f7, 0x80000, 0x800f7, 0x80014, 0x80014, 0x3f, 0x3e003f, 0x48003f, 0x24003f, 0x4003f, 0x4003f, 0x38003f, 0x3f, 0xc0022, 0x1e007f, 0x330022, 0x330014, 0x3f0073, 0x330012, 0x330072, 0x0, 0x380000, 0x0, 0x1f0000, 0x330000, 0x330000, 0x330018, 0x330030, 0x30001e, 0x240008, 0x4f0008, 0x40008, 0x3c0000, 0x460000, 0x450008, 0x220008, 0x8, 0x7f, 0x2, 0x18003e, 0x22, 0x7e0022, 0x600022, 0x7e0011, 0x0, 0x180038, 0x18003c, 0x180036, 0x180033, 0x1f007f, 0x30, 0x78, 0x0, 0x18, 0x3e0018, 0x100018, 0x1f, 0xff, 0x18, 0x18, 0x18, 0x4, 0x7e, 0x1e0049, 0x330008, 0x330008, 0x330008, 0x1e0008, 0x0, 0x400000, 0x800000, 0x200000, 0x4000ff, 0xff, 0x18, 0x18, 0x18, 0x8, 0x1e0008, 0x330008, 0x1f0008, 0x330008, 0x1f0008, 0x3000c, 0x30000, 0x7f0000, 0x10038, 0x3f0000, 0x210033, 0x3f0033, 0x10033, 0x7f007e, 0x0, 0x180000, 0x180000, 0x180000, 0x18003f, 0x180030, 0x180030, 0x180000, 0x180000, 0x18, 0x200018, 0x220018, 0x220018, 0x2a00f8, 0x240000, 0x100000, 0x0, 0xf0000, 0xf0000, 0xf001e, 0xf0033, 0xf0033, 0xf0033, 0xf001e, 0xf0000, 0x70000, 0x60000, 0x660033, 0x36007f, 0x1e007f, 0x36006b, 0x670063, 0x0, 0x8001e, 0x80033, 0x80000, 0xf0033, 0xff0033, 0x33, 0x7e, 0x0, 0x8, 0x8, 0x8, 0x6600f8, 0x8, 0x6600f8, 0x8, 0x8, 0x10067, 0x10066, 0x10036, 0x1001e, 0x10036, 0x10066, 0x10067, 0x10000, 0x1c00f0, 0x800f0, 0x3e00f0, 0x2200f0, 0x4100f0, 0x1400f0, 0x1400f0, 0xf0, 0x22, 0x2a, 0x7f002a, 0x2a, 0x3002a, 0x1c0022, 0x600021, 0x0, 0x3f0000, 0x660000, 0x660000, 0x3e0036, 0x60000, 0x6007f, 0xf0000, 0x0, 0x8, 0x8, 0xff0008, 0xff000f, 0xff0008, 0xff000f, 0xff0000, 0xff0000, 0x4f005d, 0x4800b1, 0x48005d, 0x7f0011, 0x90039, 0x90055, 0x790009, 0x0, 0x3c0008, 0x360008, 0x360008, 0x7c000f, 0xf, 0x8, 0x8, 0x8, 0x0, 0x70000, 0xfe, 0x1e0030, 0x3300fe, 0x330033, 0x1e00fe, 0x0, 0xe001e, 0xc0033, 0xc0030, 0xc001c, 0xc0006, 0xc0033, 0x1e003f, 0x0, 0x1c0008, 0x36000c, 0x36003e, 0x1c000c, 0xc, 0x2c, 0x18, 0x0, 0x38007c, 0x36, 0x1e0033, 0xc007f, 0xc0033, 0xc0033, 0x1e0073, 0x0, 0x80000, 0x1, 0x80002, 0x80004, 0x8004f, 0x90, 0x800a0, 0x80040, 0x0, 0x0, 0x3f, 0xf80000, 0xf80000, 0x3f, 0x0, 0x0, 0x0, 0x0, 0x36003c, 0x360060, 0x360030, 0x360018, 0x36000c, 0x0, 0x0, 0x0, 0x330000, 0x1b0000, 0xf00ff, 0x1b0008, 0x330008, 0x8, 0x180000, 0x180000, 0x180000, 0x1f0055, 0x1f0055, 0x80000, 0x80000, 0x80000, 0x7d0018, 0x110018, 0x7d0018, 0x110018, 0x3900ff, 0x550008, 0x90008, 0x8, 0x63000a, 0x63000a, 0x36003f, 0x1c002a, 0x36002a, 0x630002, 0x63007e, 0x0, 0x3e001e, 0x140033, 0x140000, 0x55001e, 0x550033, 0x140033, 0x7f001e, 0x0, 0x36007f, 0x49, 0x7f0049, 0x220049, 0x7f0055, 0x220063, 0x21007f, 0x0, 0x36, 0x36, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

// Font8x8 is a fixed-width 8x8 pixel font which has been made available in the Public Domain.
// It looks a little bit like this (somewhat heavy, serifs, typical console font from early PCs)
//
//              XXX                         X      XXXX              XXXX
//             XX XX                       XX     XX  XX            XX  XX
//             XX       XXXX    XXXXX     XXXXX   XX  XX   XX   XX  XX  XX
//            XXXX     XX  XX   XX  XX     XX      XXXX     XX XX    XXXX
//             XX      XX  XX   XX  XX     XX     XX  XX     XXX    XX  XX
//             XX      XX  XX   XX  XX     XX X   XX  XX    XX XX   XX  XX
//            XXXX      XXXX    XX  XX      XX     XXXX    XX   XX   XXXX
//
var Font8x8 = &PixFont{
	8, 8,
	eightMap,
	eightData,
	false,
}
