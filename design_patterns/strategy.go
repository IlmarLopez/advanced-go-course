package main

import "fmt"

type PasswordProtecter struct {
	user         string
	passwordName string
	hashAlgoritm HashAlgorithm
}

type HashAlgorithm interface {
	Hash(p *PasswordProtecter)
}

func NewPasswordProtector(user, passwordName string, hash HashAlgorithm) *PasswordProtecter {
	return &PasswordProtecter{
		user:         user,
		passwordName: passwordName,
		hashAlgoritm: hash,
	}
}

func (p *PasswordProtecter) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgoritm = hash
}

func (p *PasswordProtecter) Hash() {
	p.hashAlgoritm.Hash(p)
}

type SHA struct{}

func (SHA) Hash(p *PasswordProtecter) {
	fmt.Printf("Hashing using SHA for %s\n", p.passwordName)
}

type MD5 struct{}

func (MD5) Hash(p *PasswordProtecter) {
	fmt.Printf("Hashing using MD5 for %s\n", p.passwordName)
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	passwordProtecter := NewPasswordProtector("ilmar", "gmail password", sha)
	passwordProtecter.Hash()

	passwordProtecter.SetHashAlgorithm(md5)
	passwordProtecter.Hash()
}
