package daomock

//go:generate mockgen -destination=mock.go -package=$GOPACKAGE github.com/liluhao/ginstarter/pkg/dao MemberDAO
