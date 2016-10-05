```
~/g/s/g/d/config-demo ❯❯❯ go run *.go
Config is:
 &config.Config{
    Value1: "first value",
    Value2: "second value",
    Stages: {
        {
            "firstString": "first string",
            "firstInt":    int(1),
            "name":        "first",
        },
        {
            "name":         "second",
            "secondString": "second string",
            "secondInt":    int(12),
        },
    },
}
Real stages are:
 []stages.Stage{
    &main.FirstStage{FirstString:"first string", FirstInt:1},
    &main.SecondStage{SecondString:"second string", SecondInt:12},
}
1. Running stage "FirstStage":
I am a first stage!
2. Running stage "SecondStage":
I'm a second stage!
```
