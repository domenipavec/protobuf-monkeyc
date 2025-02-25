// Code generated by protoc-gen-monkeyc. DO NOT EDIT.

import Toybox.Lang;

class ExampleMessage {
    class LocalMessage {
        public var l1 as String;
        
        public function initialize() {
            l1 = "";
        }
    
        public function Encode() as ByteArray {
            var result = []b;
            result.addAll(Protobuf.encodeFieldLen(1, l1, false));
            return result;
        }
        public function Decode(input as ByteArray) as Void {
            var d = new Protobuf.Decoder(input);
            while (d.remaining() > 0) {
                var tag = d.varint32();
                switch (tag >> 3) {
                    case 1: {
                        Protobuf.assertWireType(tag, Protobuf.LEN);
                        l1 = d.string();
                        break;
                    }
                }
            }
        }
    }

    enum LocalEnum {
        LA = 0,
        LB = 1,
    }

    public var i32 as Number;
    public var i64 as Long;
    public var u32 as Number;
    public var u64 as Long;
    public var s32 as Number;
    public var s64 as Long;
    public var f32 as Number;
    public var f64 as Long;
    public var sf32 as Number;
    public var sf64 as Long;
    public var fl as Float;
    public var str as String;
    public var byt as ByteArray;
    public var b as Boolean;
    public var ge as GlobalEnum;
    public var le as LocalEnum;
    public var gm as GlobalMessage;
    public var lm as LocalMessage;
    public var ri64 as Array<Long>;
    public var rf32 as Array<Number>;
    public var rf64 as Array<Long>;
    public var rstr as Array<String>;
    public var rgm as Array<GlobalMessage>;
    public var rpi64 as Array<Long>;
    public var rpf32 as Array<Number>;
    public var rpf64 as Array<Long>;
    
    public function initialize() {
        i32 = 0;
        i64 = 0l;
        u32 = 0;
        u64 = 0l;
        s32 = 0;
        s64 = 0l;
        f32 = 0;
        f64 = 0l;
        sf32 = 0;
        sf64 = 0l;
        fl = 0.0;
        str = "";
        byt = []b;
        b = false;
        ge = 0 as GlobalEnum;
        le = 0 as LocalEnum;
        gm = new GlobalMessage();
        lm = new LocalMessage();
        ri64 = [];
        rf32 = [];
        rf64 = [];
        rstr = [];
        rgm = [];
        rpi64 = [];
        rpf32 = [];
        rpf64 = [];
    }

    public function Encode() as ByteArray {
        var result = []b;
        result.addAll(Protobuf.encodeFieldVarint(1, i32, false));
        result.addAll(Protobuf.encodeFieldVarint(2, i64, false));
        result.addAll(Protobuf.encodeFieldVarint(3, u32, false));
        result.addAll(Protobuf.encodeFieldVarint(4, u64, false));
        result.addAll(Protobuf.encodeFieldVarint(5, Protobuf.toSignedInt(s32), false));
        result.addAll(Protobuf.encodeFieldVarint(6, Protobuf.toSignedInt(s64), false));
        result.addAll(Protobuf.encodeField32(7, f32, false));
        result.addAll(Protobuf.encodeField64(8, f64, false));
        result.addAll(Protobuf.encodeField32(9, sf32, false));
        result.addAll(Protobuf.encodeField64(10, sf64, false));
        result.addAll(Protobuf.encodeField32(11, fl, false));
        result.addAll(Protobuf.encodeFieldLen(13, str, false));
        result.addAll(Protobuf.encodeFieldLen(14, byt, false));
        result.addAll(Protobuf.encodeFieldVarint(15, b, false));
        result.addAll(Protobuf.encodeFieldVarint(16, ge, false));
        result.addAll(Protobuf.encodeFieldVarint(17, le, false));
        result.addAll(Protobuf.encodeFieldLen(18, gm.Encode(), false));
        result.addAll(Protobuf.encodeFieldLen(19, lm.Encode(), false));
        for (var i = 0; i < ri64.size(); i++) {
            result.addAll(Protobuf.encodeFieldVarint(20, ri64[i], true));
        }
        for (var i = 0; i < rf32.size(); i++) {
            result.addAll(Protobuf.encodeField32(21, rf32[i], true));
        }
        for (var i = 0; i < rf64.size(); i++) {
            result.addAll(Protobuf.encodeField64(22, rf64[i], true));
        }
        for (var i = 0; i < rstr.size(); i++) {
            result.addAll(Protobuf.encodeFieldLen(23, rstr[i], true));
        }
        for (var i = 0; i < rgm.size(); i++) {
            result.addAll(Protobuf.encodeFieldLen(24, rgm[i].Encode(), true));
        }
        {
            var packed = []b;
            for (var i = 0; i < rpi64.size(); i++) {
                packed.addAll(Protobuf.encodeFieldVarint(0, rpi64[i], true));
            }
            result.addAll(Protobuf.encodeFieldLen(25, packed, false));
        }
        {
            var packed = []b;
            for (var i = 0; i < rpf32.size(); i++) {
                packed.addAll(Protobuf.encodeField32(0, rpf32[i], true));
            }
            result.addAll(Protobuf.encodeFieldLen(26, packed, false));
        }
        {
            var packed = []b;
            for (var i = 0; i < rpf64.size(); i++) {
                packed.addAll(Protobuf.encodeField64(0, rpf64[i], true));
            }
            result.addAll(Protobuf.encodeFieldLen(27, packed, false));
        }
        return result;
    }
    public function Decode(input as ByteArray) as Void {
        var d = new Protobuf.Decoder(input);
        while (d.remaining() > 0) {
            var tag = d.varint32();
            switch (tag >> 3) {
                case 1: {
                    Protobuf.assertWireType(tag, Protobuf.VARINT);
                    i32 = d.varint32();
                    break;
                }
                case 2: {
                    Protobuf.assertWireType(tag, Protobuf.VARINT);
                    i64 = d.varint64();
                    break;
                }
                case 3: {
                    Protobuf.assertWireType(tag, Protobuf.VARINT);
                    u32 = d.varint32();
                    break;
                }
                case 4: {
                    Protobuf.assertWireType(tag, Protobuf.VARINT);
                    u64 = d.varint64();
                    break;
                }
                case 5: {
                    Protobuf.assertWireType(tag, Protobuf.VARINT);
                    s32 = Protobuf.fromSignedNumber(d.varint32());
                    break;
                }
                case 6: {
                    Protobuf.assertWireType(tag, Protobuf.VARINT);
                    s64 = Protobuf.fromSignedLong(d.varint64());
                    break;
                }
                case 7: {
                    Protobuf.assertWireType(tag, Protobuf.I32);
                    f32 = d.number();
                    break;
                }
                case 8: {
                    Protobuf.assertWireType(tag, Protobuf.I64);
                    f64 = d.long();
                    break;
                }
                case 9: {
                    Protobuf.assertWireType(tag, Protobuf.I32);
                    sf32 = d.number();
                    break;
                }
                case 10: {
                    Protobuf.assertWireType(tag, Protobuf.I64);
                    sf64 = d.long();
                    break;
                }
                case 11: {
                    Protobuf.assertWireType(tag, Protobuf.I32);
                    fl = d.float();
                    break;
                }
                case 13: {
                    Protobuf.assertWireType(tag, Protobuf.LEN);
                    str = d.string();
                    break;
                }
                case 14: {
                    Protobuf.assertWireType(tag, Protobuf.LEN);
                    byt = d.data();
                    break;
                }
                case 15: {
                    Protobuf.assertWireType(tag, Protobuf.VARINT);
                    b = d.varint32() != 0;
                    break;
                }
                case 16: {
                    Protobuf.assertWireType(tag, Protobuf.VARINT);
                    ge = d.varint32() as GlobalEnum;
                    break;
                }
                case 17: {
                    Protobuf.assertWireType(tag, Protobuf.VARINT);
                    le = d.varint32() as LocalEnum;
                    break;
                }
                case 18: {
                    Protobuf.assertWireType(tag, Protobuf.LEN);
                    gm.Decode(d.data());
                    break;
                }
                case 19: {
                    Protobuf.assertWireType(tag, Protobuf.LEN);
                    lm.Decode(d.data());
                    break;
                }
                case 20: {
                    switch (tag & 7) {
                        case Protobuf.VARINT:
                            ri64.add(d.varint64());
                            break;
                        case Protobuf.LEN:
                            for (var endRemaining = d.remaining() - d.varint32(); d.remaining() > endRemaining;) {
                                ri64.add(d.varint64());
                            }
                            break;
                        default:
                            throw new Protobuf.Exception("invalid wire type");
                    }
                    break;
                }
                case 21: {
                    switch (tag & 7) {
                        case Protobuf.I32:
                            rf32.add(d.number());
                            break;
                        case Protobuf.LEN:
                            for (var endRemaining = d.remaining() - d.varint32(); d.remaining() > endRemaining;) {
                                rf32.add(d.number());
                            }
                            break;
                        default:
                            throw new Protobuf.Exception("invalid wire type");
                    }
                    break;
                }
                case 22: {
                    switch (tag & 7) {
                        case Protobuf.I64:
                            rf64.add(d.long());
                            break;
                        case Protobuf.LEN:
                            for (var endRemaining = d.remaining() - d.varint32(); d.remaining() > endRemaining;) {
                                rf64.add(d.long());
                            }
                            break;
                        default:
                            throw new Protobuf.Exception("invalid wire type");
                    }
                    break;
                }
                case 23: {
                    Protobuf.assertWireType(tag, Protobuf.LEN);
                    rstr.add(d.string());
                    break;
                }
                case 24: {
                    Protobuf.assertWireType(tag, Protobuf.LEN);
                    var msg = new GlobalMessage();
                    msg.Decode(d.data());
                    rgm.add(msg);
                    break;
                }
                case 25: {
                    switch (tag & 7) {
                        case Protobuf.VARINT:
                            rpi64.add(d.varint64());
                            break;
                        case Protobuf.LEN:
                            for (var endRemaining = d.remaining() - d.varint32(); d.remaining() > endRemaining;) {
                                rpi64.add(d.varint64());
                            }
                            break;
                        default:
                            throw new Protobuf.Exception("invalid wire type");
                    }
                    break;
                }
                case 26: {
                    switch (tag & 7) {
                        case Protobuf.I32:
                            rpf32.add(d.number());
                            break;
                        case Protobuf.LEN:
                            for (var endRemaining = d.remaining() - d.varint32(); d.remaining() > endRemaining;) {
                                rpf32.add(d.number());
                            }
                            break;
                        default:
                            throw new Protobuf.Exception("invalid wire type");
                    }
                    break;
                }
                case 27: {
                    switch (tag & 7) {
                        case Protobuf.I64:
                            rpf64.add(d.long());
                            break;
                        case Protobuf.LEN:
                            for (var endRemaining = d.remaining() - d.varint32(); d.remaining() > endRemaining;) {
                                rpf64.add(d.long());
                            }
                            break;
                        default:
                            throw new Protobuf.Exception("invalid wire type");
                    }
                    break;
                }
            }
        }
    }
}

class GlobalMessage {
    public var g1 as Number;
    
    public function initialize() {
        g1 = 0;
    }

    public function Encode() as ByteArray {
        var result = []b;
        result.addAll(Protobuf.encodeFieldVarint(1, g1, false));
        return result;
    }
    public function Decode(input as ByteArray) as Void {
        var d = new Protobuf.Decoder(input);
        while (d.remaining() > 0) {
            var tag = d.varint32();
            switch (tag >> 3) {
                case 1: {
                    Protobuf.assertWireType(tag, Protobuf.VARINT);
                    g1 = d.varint32();
                    break;
                }
            }
        }
    }
}

enum GlobalEnum {
    A = 0,
    B = 1,
}
