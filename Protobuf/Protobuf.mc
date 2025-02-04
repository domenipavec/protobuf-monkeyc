import Toybox.Lang;
using Toybox.StringUtil;

module Protobuf {
    enum WireType {
        VARINT = 0,
        I64 = 1,
        LEN = 2,
        SGROUP = 3,
        EGROUP = 4,
        I32 = 5,
    }

    function encodeFieldVarint(f as Number, v as Number or Long or Boolean, force as Boolean) as ByteArray {
        if (((v instanceof Boolean && !v) || v == 0) && !force) {
            return []b;
        }
        var result = []b;
        result.addAll(encodeTag(f, VARINT));
        result.addAll(encodeVarint(v));
        return result;
    }

    function encodeField32(f as Number, v as Number or Float, force as Boolean) as ByteArray {
        if (v == 0 && !force) {
            return []b;
        }
        var result = []b;
        result.addAll(encodeTag(f, I32));
        var format = Lang.NUMBER_FORMAT_SINT32;
        if (v instanceof Float) {
            format = Lang.NUMBER_FORMAT_FLOAT;
        }
        result.addAll((new [4]b).encodeNumber(v, format, {:endianness => Lang.ENDIAN_LITTLE}));
        return result;
    }

    function encodeField64(f as Number, v as Long, force as Boolean) as ByteArray {
        if (v == 0 && !force) {
            return []b;
        }
        var result = []b;
        result.addAll(encodeTag(f, I64));
        result.addAll((new [4]b).encodeNumber(v&0xffffffff, Lang.NUMBER_FORMAT_SINT32, {:endianness => Lang.ENDIAN_LITTLE}));
        result.addAll((new [4]b).encodeNumber(v>>32, Lang.NUMBER_FORMAT_SINT32, {:endianness => Lang.ENDIAN_LITTLE}));
        return result;
    }

    function encodeFieldLen(f as Number, v as String or ByteArray, force as Boolean) as ByteArray {
        if (v instanceof String) {
            v = StringUtil.convertEncodedString(v, {
                :fromRepresentation => StringUtil.REPRESENTATION_STRING_PLAIN_TEXT,
                :toRepresentation => StringUtil.REPRESENTATION_BYTE_ARRAY,
                :encoding => StringUtil.CHAR_ENCODING_UTF8,
            }) as ByteArray;
        }
        if (v.size() == 0 && !force) {
            return []b;
        }
        var result = []b;
        result.addAll(encodeTag(f, LEN));
        result.addAll(encodeVarint(v.size()));
        result.addAll(v);
        return result;
    }

    function encodeTag(f as Number, w as WireType) as ByteArray {
        if (f <= 0) {
            return []b;
        }
        return encodeVarint(f << 3 | w);
    }

    function toSignedInt(v as Number or Long) as Number or Long {
        if (v instanceof Number) {
            return (v << 1) ^ (v >> 31);
        } else {
            return (v << 1) ^ (v >> 63);
        }
    }

    function encodeVarint(v as Number or Long or Boolean) as ByteArray {
        if (v instanceof Boolean) {
            v = v ? 1 : 0;
        }
        var result = []b;
        do {
            var b = (v & 0x7f).toNumber();
            v = v >> 7;
            // remove negative bit that replicates on shift
            v &= 0x01ffffffffffffffl;
            if (v > 0) {
                b |= (1<<7);
            }
            result.add(b);
        } while (v != 0);
        return result;
    }
}
