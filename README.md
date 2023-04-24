# Terraform Provider for Cattle Names

This repository contains a Terraform provider that generates unique names based on different families, such as cattle names, Game of Thrones characters, and members of the 27 Club. This project is intended for learning purposes and is not designed for real-world use.

## Features

- Generate random names from predefined families:
  - `cattle`: Names like "Bessie" and "MooMoo".
  - `GOT`: Names of Game of Thrones characters like "Jon Snow" and "Daenerys".
  - `27Club`: Names of 27 Club members like "Kurt Cobain" and "Amy Winehouse".

## Repository Structure

- `docs/`: Contains documentation for the provider.
- `examples/`: Includes example Terraform configurations to demonstrate usage.
- `internal/`: Contains the provider and resource implementation.
  - `provider/`: Defines the provider.
  - `resource/`: Implements the resource logic.

## Usage

This provider is not published to the Terraform Registry. To use it, clone this repository and build the provider locally. You can then use the examples provided in the `examples/` directory to test its functionality.

## Disclaimer

This project was created for educational purposes only. It is not intended for production use and may not follow best practices for Terraform provider development. Use it at your own risk.