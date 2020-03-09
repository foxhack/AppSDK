package nrprivate

import (
	"net/url"
)



// Printing function purely for debugging purposes
// Print the body of a request to the console

func ishex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}


// unescape unescapes a string; the mode specifies
// which section of the URL string is being unescaped.
func edgexunescape(s string) (string, error) {
	// Count %, check that they're well-formed.
	n := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case '$':
			n++
			if i+2 >= len(s) || !ishex(s[i+1]) || !ishex(s[i+2]) {
				s = s[i:]
				if len(s) > 3 {
					s = s[:3]
				}
				return "", url.EscapeError(s)
			}
			i=i+3
		default:

			i++
		}
	}

	if n == 0  {
		return s, nil
	}

	t := make([]byte, len(s)-2*n)
	j := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case '$':
			t[j] = unhex(s[i+1])<<4 | unhex(s[i+2])
			j++
			i += 3
		default:
			t[j] = s[i]
			j++
			i++
		}
	}
	return string(t), nil
}

func nrdevicenameunescape( device string ,name string ) (string , string ,error ){

	deviceId, err := url.QueryUnescape(device)

	// Problems unescaping URL
	if err != nil {
		//LoggingClient.Error("Error unescaping the device  ID: " + err.Error())
		return deviceId,name ,err
	}

	name, err = url.QueryUnescape(name)
	// Problems unescaping URL
	if err != nil {

		//LoggingClient.Error("Error unescaping the name  ID: " + err.Error())
		return deviceId,name ,err
	}

	deviceId,err=edgexunescape(deviceId)
	if err != nil {
		//LoggingClient.Error("Error edgexunescape the device ID: " + err.Error())
		return deviceId,name ,err
	}
	name,err=edgexunescape(name)
	if err != nil {

		//LoggingClient.Error("Error edgexunescape the name  ID: " + err.Error())
		return deviceId,name ,err
	}
	//LoggingClient.Error(fmt.Sprint(name))
	return deviceId,name ,nil
}


func NrNameUnescape( name string ) (string ,error ){
	//LoggingClient.Error(fmt.Sprint(name))
	name, err := url.QueryUnescape(name)
	// Problems unescaping URL
	if err != nil {
		//LoggingClient.Error("Error  nrnameunescape the name  ID: " + err.Error())
		return name ,err
	}
	name,err=edgexunescape(name)
	if err != nil {

	//	LoggingClient.Error("Error edgexunescape the name  ID: " + err.Error())
		return name ,err
	}
	//LoggingClient.Error(fmt.Sprint(name))
	return name ,nil
}



func NrEscape(s string ) string {
	spaceCount, hexCount := 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			if c == ' '{
				spaceCount++
			} else {
				hexCount++
			}
		}
	}

	if spaceCount == 0 && hexCount == 0 {
		return s
	}

	var buf [64]byte
	var t []byte

	required := len(s) + 2*hexCount
	if required <= len(buf) {
		t = buf[:required]
	} else {
		t = make([]byte, required)
	}

	if hexCount == 0 {
		copy(t, s)
		for i := 0; i < len(s); i++ {
			if s[i] == ' ' {
				t[i] = '+'
			}
		}
		return string(t)
	}

	j := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case shouldEscape(c):
			t[j] = '$'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		default:
			t[j] = s[i]
			j++
		}
	}
	return string(t)
}



func shouldEscape(c byte) bool {
	// §2.3 Unreserved characters (alphanum)
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}

	switch c {
	case '$', '&', '+', ',', '/', ':', ';', '=', '?', '@': // §2.2 Reserved characters (reserved)
		// Different sections of the URL allow a few of
		// the reserved characters to appear unescaped.

		return true
	}


	// Everything else must be escaped.
	return false
}




func NrUnescape(s string) (string, error) {
	// Count %, check that they're well-formed.
	n := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case '$':
			n++
			if i+2 >= len(s) || !ishex(s[i+1]) || !ishex(s[i+2]) {
				s = s[i:]
				if len(s) > 3 {
					s = s[:3]
				}
				return "", url.EscapeError(s)
			}
			i=i+3
		default:

			i++
		}
	}

	if n == 0  {
		return s, nil
	}

	t := make([]byte, len(s)-2*n)
	j := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case '$':
			t[j] = unhex(s[i+1])<<4 | unhex(s[i+2])
			j++
			i += 3
		default:
			t[j] = s[i]
			j++
			i++
		}
	}
	return string(t), nil
}

