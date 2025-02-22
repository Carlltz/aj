# aj

**Currently, `aj` is only available for the Fish shell, but you are more than welcome to add support for your preferred shell!**

`aj` is a simple shell tool that automatically fixes mistakes in your last command. If you mistype a command, just type `aj`, and it will suggest a correction based on what you likely meant to run.

## Fun Fact

Did you know that "aj" means "ouch" in Swedish? That's why this tool is named `aj` — to help ease the pain of mistyped commands! Plus, `aj` is quite similar to `ai`, which powers the tool to correct your commands.

## Features

-   Automatically detects and fixes common typos in shell commands
-   Suggests a corrected command and prompts before execution
-   Simple and lightweight, built with Go

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Carlltz/aj
    cd aj
    ```

2. Create a `.env` file based on `.env.example` and add your OpenAI API key:

    ```sh
    cp .env.example .env
    ```

    Then, edit the `.env` file to include your OpenAI API key:

    ```sh
    OPENAI_API_KEY=your-api-key-here
    ```

3. Build the project:

    ```sh
    go build -o aj
    ```

4. Move the binary to a directory in your `$PATH` (optional but recommended):

    ```sh
    sudo mv aj /usr/local/bin/
    ```

## Usage

Simply type `aj` after running a command that failed:

```sh
$ eco Hello, World!
$ aj

Last command: eco Hello, World!

Ctrl+C to exit or Enter to run: echo Hello, World!
```

Press **Enter** to run the suggested command or **Ctrl+C** to cancel.

## Cost

The usage of the OpenAI API is very affordable. Approximately every 15-25 corrected commands cost about 0.001 USD, making it a cost-effective solution for fixing command-line mistakes.

## Configuration

The `.env` file must contain a valid OpenAI API key to function properly. You can obtain an API key from [OpenAI](https://openai.com/api/).

## Disclaimer

I'm not very familiar with Go, so I might not have followed "best practices" at all. If you have suggestions for improvement, feel free to contribute!

## License

MIT License

## Contributing

Pull requests are welcome! Feel free to open an issue if you encounter any bugs or have feature suggestions.
