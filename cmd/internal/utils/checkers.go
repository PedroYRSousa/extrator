package utils

func CheckIsHardCodedSecretOrEnv(target string) bool {
	return IsHardcoded(target) || IsEnv(target) || IsSecret(target)
}
