package all

import (
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/buildinfo"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/command/apk"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/config/all"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/executable"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/c/conan"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/dotnet/deps"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/dotnet/nuget"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/golang/binary"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/golang/mod"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/java/gradle"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/java/jar"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/java/pom"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/nodejs/npm"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/nodejs/pkg"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/nodejs/pnpm"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/nodejs/yarn"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/php/composer"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/python/packaging"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/python/pip"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/python/pipenv"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/python/poetry"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/ruby/bundler"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/ruby/gemspec"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/rust/binary"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language/rust/cargo"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/licensing"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/os/alpine"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/os/amazonlinux"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/os/debian"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/os/mariner"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/os/redhatbase"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/os/release"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/os/ubuntu"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/pkg/apk"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/pkg/dpkg"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/pkg/rpm"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/repo/apk"
	_ "github.com/khulnasoft/tunnel/pkg/fanal/analyzer/secret"
)
