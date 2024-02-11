# Population Generation Project

## Overview

This project is designed to simulate the generation of a virtual population, providing a detailed breakdown of demographic attributes such as age, sex, personality type, and other characteristics. The simulation operates in two primary phases: first, it calculates a proportional breakdown of the population based on input criteria; and second, it generates individual profiles for each member of the population, tracking the distribution of attributes to ensure adherence to the initially calculated proportions.

## Features

- **Population Breakdown**: Generates a YAML file outlining the proportional distribution of various attributes (e.g., sex, personality type) based on the total population size specified by the user.
- **Individual Profile Generation**: Iteratively creates detailed YAML profiles for each population member, decrementing from the attribute quotas as profiles are generated to maintain the defined distribution.
- **Attribute Tracking**: Monitors the allocation of attributes throughout the profile generation process to ensure the final population composition matches the initial breakdown.

## Getting Started

### Prerequisites

- Go programming environment (1.15 or later recommended)
- Basic understanding of YAML structure and Go's `yaml` package for marshaling/unmarshaling data

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/yourgithubusername/population-gen.git
cd population-gen
```

### Usage

1. **Generate Population Breakdown**:

   Define your population criteria in a configuration file or directly within the script to generate the initial population breakdown as a YAML file.

   ```go
   go run main.go --breakdown
   ```

2. **Generate Individual Profiles**:

   Once the population breakdown is defined, run the profile generation script to create individual YAML files for each member of the population.

   ```go
   go run main.go --generate-profiles
   ```

## Configuration

Customize your population generation by modifying the input criteria. This includes defining the total population size and the proportions for each attribute (e.g., 50% male, 50% female). These settings can be adjusted in the `config.yaml` file or directly within the source code, depending on your implementation.

## Contributing

Contributions to the Population Generation Project are welcome! Please refer to the contributing guidelines for more details on how to submit pull requests, report issues, or request features.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for providing extensive documentation and libraries that made this project possible.
- Special thanks to contributors and users who have tested and provided feedback for this tool.
