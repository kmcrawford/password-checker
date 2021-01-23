package main

import (
  "bufio"
  "crypto/sha1"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "strconv"
  "strings"
)

func main() {
  if len(os.Args) == 1 {
    fmt.Println("Usage: password-checker <password-to-check>")
    os.Exit(2)
  }

  pw := os.Args[1]

  pwPrefix, pwSuffix := hashPassword(pw)

  err, hashedSuffixes := retrievePasswordHashesFromApi(pwPrefix)
  if err != nil {
    panic(err)
  }
  err, _ = checkForCompromisedPassword(hashedSuffixes, pwSuffix)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  fmt.Println("Password not found to be compromised")
}

func checkForCompromisedPassword(suffixResponse string, pwSuffix string) (error, int) {
  scanner := bufio.NewScanner(strings.NewReader(suffixResponse))
  for scanner.Scan() {
    i := scanner.Text()
    d := strings.Split(i, ":")
    if d[1] == "0" {
      continue
    }
    if d[0] == pwSuffix {
      compromiseCount, _ := strconv.Atoi(d[1])
      return errors.New("Appeared "+ d[1]+ " times in security breaches"), compromiseCount
    }
  }
  return nil, 0
}

func retrievePasswordHashesFromApi(pwPrefix string) (err error, hashedSuffixes string) {
  url := "https://api.pwnedpasswords.com/range/" + pwPrefix
  method := "GET"
  client := &http.Client{}
  req, err := http.NewRequest(method, url, nil)
  req.Header.Add("Add-Padding", "true")
  if err != nil {
    fmt.Println(err)
    return
  }
  res, err := client.Do(req)
  if err != nil {
    return
  }
  defer res.Body.Close()
  b, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return
  }
  hashedSuffixes = string(b)
  return
}

func hashPassword(pw string) (string, string) {
  h := sha1.New()
  h.Write([]byte(pw))
  bs := h.Sum(nil)
  hashPw := fmt.Sprintf("%x", bs)
  return hashPw[:5], strings.ToUpper(hashPw[5:])
}

