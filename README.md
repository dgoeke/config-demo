```
~/g/s/g/d/config-demo ❯❯❯ go run *.go                                                                                             V master
Config is:
 &config.Config{
    Value1: "first value",
    Value2: "second value",
    Stages: {
        {
            "name":        "first",
            "firstString": "first string",
            "firstInt":    int(1),
        },
        {
            "name":         "second",
            "secondString": "second string",
            "secondInt":    int(12),
        },
    },
    realStages: {
        &main.FirstStage{FirstString:"first string", FirstInt:1},
        &main.SecondStage{SecondString:"second string", SecondInt:12},
    },
}
0. Running stage "FirstStage":
I am a first stage!
1. Running stage "SecondStage":
I'm a second stage!
```
