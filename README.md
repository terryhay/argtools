# argtools
The module for work with command line arguments.

The main advantage *argtools* is printing perfect formatted help information and checking input argument data.

## Setup

Import *argtools* module into your project:

`go get github.com/terryhay/argtools`

Build *argtools* generator into some directory where you are keeping bin files for your project:

`go build -o ./bin/gen_argtools github.com/terryhay/argtools/internal/generator`

The *argtools* generator will be built into *./bin* directory and will be named *gen_argtools*.

OK, the next step is to write *arg_config.yaml* file with description of your application commands, flags and other data for help info. For more information (you need it I guess) you can research arg_config.yaml files in example directories of argtools module. Good luck.

If your *arg_tools_config.yaml* is in *./config* directory you can generate

`./bin/gen_argtools -c ./config/arg_tools_config.yaml -o ./`

The *argtools* generator will create *./argTools* directory with *arg_tools.go* file. The file will contain *Parse* method.

Besides you can add strings

`go build -o ./bin/gen_argtools github.com/terryhay/argtools/internal/generator`

`./bin/gen_argtools -c ./config/arg_tools_config.yaml -o ./`

into your *Makefile* and use `make generate` command like it is done in examples.