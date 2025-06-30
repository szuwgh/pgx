package smx509

// Possible certificate files; stop after finding one.
var certFiles = []string{
	"/etc/certs/ca-certificates.crt",     // Solaris 11.2+
	"/etc/ssl/certs/ca-certificates.crt", // Joyent SmartOS
	"/etc/ssl/cacert.pem",                // OmniOS
}

// Possible directories with certificate files; all will be read.
var certDirectories = []string{
	"/etc/certs/CA",
}
