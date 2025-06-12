module github.com/andrewbenton/go-secrets/mapping/mapstructurev2

go 1.24.3

replace github.com/andrewbenton/go-secrets => ../..

require (
	github.com/andrewbenton/go-secrets v0.0.5
	github.com/go-viper/mapstructure/v2 v2.2.1
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
