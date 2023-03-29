package provider

import (
	"fmt"
)

func validatePrivateASN(val interface{}, key string) (warns []string, errs []error) {
	v := val.(int)
	if (v >= 64512 && v <= 65534) || (v >= 4200000000 && v <= 4294967294) {
		return
	}
	errs = append(errs, fmt.Errorf("%q must be in the range 64512 - 65534 or 4200000000 - 4294967294, got: %d", key, v))
	return
}

func validatePublicOrPrivateASN(val interface{}, key string) (warns []string, errs []error) {
	v := val.(int)
	if (v >= 1 && v <= 2147483647) || (v >= 64512 && v <= 65534) {
		return
	}
	errs = append(errs, fmt.Errorf("%q must be a public ASN (1 - 2147483647) or a private ASN (64512 - 65534), got: %d", key, v))
	return
}
