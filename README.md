<h1>Password Checker</h1>

This is an example program that uses https://haveibeenpwned.com/Passwords api to check to see if a password has been previously compromised in a data breach.

Range api used (with padding):
https://haveibeenpwned.com/API/v3#SearchingPwnedPasswordsByRange

https://www.troyhunt.com/enhancing-pwned-passwords-privacy-with-padding/

<h3>Usage:</h3>
```
->go get github.com/kmcrawford/password-checker

->password-checker ilovepuppies
Appeared 480 times in security breaches
```

The exit code will be `0` if the password has not been included in any security breaches.  If the password has been found in a previous security breach the exit code will be `1`