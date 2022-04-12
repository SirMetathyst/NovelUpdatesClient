# NovelUpdatesClient

### Contents

- Client Cli
  - Dependencies
  - Building
  - Commands
    - Search

### Dependencies

- *Go 1.18+ (Compile Time)
- Make (Compile Time)

### Building

```git clone https://github.com/SirMetathyst/NovelUpdatesClient && cd ./NovelUpdatesClient```

```make build```

```./cmd/nu/nu --help```

### Commands

```
NAME:
nu - NovelUpdate Client

USAGE:
nu [global options] command [command options] [arguments...]

VERSION:
v0.1

COMMANDS:
search   
help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
--help, -h     show help (default: false)
--version, -v  print the version (default: false)
```

### Command/Search

```
NAME:
   nu search - 

USAGE:
   nu search [command options] [arguments...]

OPTIONS:
   --novel-type value, --nt value                  (accepts multiple inputs)
   --language value, -l value                      (accepts multiple inputs)
   --chapters-scale value, --cs value              min|max (default: decided by remote)
   --chapters value, -c value                      (default: 0)
   --release-frequency-scale value, --rfs value    min|max (default: decided by remote)
   --release-frequency value, --rf value           (default: 0)
   --reviews-scale value, --rs value               min|max (default: decided by remote)
   --reviews value, -r value                       (default: 0)
   --rating-scale value, --rts value               min|max (default: decided by remote)
   --rating value, --rt value                      (default: 0)
   --number-of-ratings-scale value, --nrts value   min|max (default: decided by remote)
   --number-of-ratings value, --nrt value          (default: 0)
   --readers-scale value, --rds value              min|max (default: decided by remote)
   --readers value, --rd value                     (default: 0)
   --first-release-date-scale value, --frds value  min|max (default: decided by remote)
   --first-release-date value, --frd value         
   --last-release-date-scale value, --lrds value   min|max (default: decided by remote)
   --last-release-date value, --lrd value          
   --genre-operator value, --gop value             and|or (default: decided by remote)
   --genre-include-name value, --gin value         (accepts multiple inputs)
   --genre-exclude-name value, --gen value         (accepts multiple inputs)
   --genre-include-id value, --gii value           (accepts multiple inputs)
   --genre-exclude-id value, --gei value           (accepts multiple inputs)
   --tag-operator value, --top value               and|or (default: decided by remote)
   --tag-include-name value, --tin value           (accepts multiple inputs)
   --tag-exclude-name value, --ten value           (accepts multiple inputs)
   --tag-include-id value, --tii value             (accepts multiple inputs)
   --tag-exclude-id value, --tei value             (accepts multiple inputs)
   --story-status value, --ss value                all|completed|ongoing|hiatus (default: decided by remote)
   --sort-by value, --sb value                     chapters|frequency|rank|rating|readers|reviews|title|last-updated (default: decided by remote)
   --order-by value, --ob value                    asc|desc (default: decided by remote)
   --page value, -p value                          (default: 1)
   --output-exclude-id, --oes                      true|false (default: false)
   --output-exclude-url, --oesu                    true|false (default: false)
   --output-exclude-title, --oet                   true|false (default: false)
   --output-exclude-chapters, --oec                true|false (default: false)
   --output-exclude-release-frequency, --oerf      true|false (default: false)
   --output-exclude-readers, --oerd                true|false (default: false)
   --output-exclude-reviews, --oer                 true|false (default: false)
   --output-exclude-last-updated, --oelu           true|false (default: false)
   --output-exclude-genre, --oeg                   true|false (default: false)
   --output-exclude-description, --oed             true|false (default: false)
   --output-exclude-short-language, --oesl         true|false (default: false)
   --output-exclude-rating, --oert                 true|false (default: false)
   --output-json, --oj                             true|false (default: false)
   --output-json-indent, --oji                     true|false (default: false)
   --help, -h                                      show help (default: false)
```