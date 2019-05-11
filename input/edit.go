package input

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

//Edit spans an editor to modify the input text and feedbacks
//the result.
func Edit(data []byte, cmdEditor []string) ([]byte, error) {
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, err
	}
	defer func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}()

	n, err := tmpfile.Write(data)
	if err != nil {
		return nil, err
	}
	if n < len(data) {
		return nil, io.ErrShortWrite
	}
	tmpfile.Close()

	cmdArgs := append(cmdEditor, tmpfile.Name())
	ed := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	ed.Stdout = os.Stdout
	ed.Stdin = os.Stdin
	ed.Stderr = os.Stderr
	err = ed.Run()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		return nil, err
	}

	return body, nil
}

//EditAsJson fires-up an editor to modify th eprovided interface
//using its JSON form.
func EditAsJson(v interface{}, cmdEditor []string) (interface{}, error) {
	j, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}

	buf, err := Edit(j, cmdEditor)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buf, &v)
	if err != nil {
		return nil, err
	}

	//I'm not that sure that v needs to be reurned as for most
	//of the cases the Unmashal directvie should have already
	//propagated the mods. It happen, that it is not working at
	//least for map (that should nee to be reallocated), so
	//result is also returned to the user.
	//
	//TODO(pirmd): it is probably not the right way to do, try
	//harder to find a correct approach
	return v, nil
}
