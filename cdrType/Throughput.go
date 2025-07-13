package cdrType

type Throughput struct { /* Sequence Type */
	GuaranteedThpt Bitrate `ber:"tagNum:0"`
	MaximumThpt    Bitrate `ber:"tagNum:1"`
}
