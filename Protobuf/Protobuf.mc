import Toybox.Lang;

module Protobuf {
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
