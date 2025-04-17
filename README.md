# SolviumPayments

SolviumPayments is a payment module created by @kleeedolinux to simplify integrating various payment systems into your Go projects with a single, easy-to-use tool. It currently supports Efi Payments' Pix API.

## Why I Created SolviumPayments

As a developer, I often found it frustrating to deal with the complexity of integrating multiple payment systems into projects. Each provider has its own API, authentication methods, and quirks, which can slow down development and make code messy. I wanted a cleaner, more unified way to handle payments without reinventing the wheel every time.

SolviumPayments was born to solve this. My goal was to build a modular, developer-friendly tool that abstracts the nitty-gritty details of payment APIs, letting you focus on your project instead of wrestling with documentation or cryptic error messages. Starting with Efi Payments' Pix API, I aimed to create a foundation that’s simple to use but flexible enough to grow with new features and providers in the future.

Whether you’re handling instant Pix charges, managing refunds, or splitting payments across accounts, SolviumPayments is designed to make your life easier with a clean API and straightforward setup. It’s all about saving time and reducing headaches for developers like me who just want payments to work.

## But Isn’t There an Official SDK?

Yes — [there is an official Go SDK provided by Efi](https://github.com/efipay/sdk-go-apis-efi/tree/main). However, in my opinion, it’s overly verbose, hard to follow, and not very idiomatic to Go. It’s essentially a thin wrapper over raw API requests using generic `map[string]interface{}` payloads, which can lead to confusing, error-prone code.

That’s one of the main reasons I created SolviumPayments.

I expect this pattern to hold true for many other payment systems I’ll integrate in the future. Rather than bending my code to fit inconsistent or poorly designed SDKs, I’m building a unified abstraction that feels natural in Go, is easier to maintain, and grows with my needs.

## Installation

```bash
go get github.com/solviumdream/solviumpayments
```

## License

AGPL-3.0

## Support

For issues or feature requests, please open an issue on GitHub.