package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    config, err := LoadConfig("config.gcfg")

    fmt.Println(config.Database.Connection)

    if err != nil {
        panic(err)
    }

    db_conn, err := NewDamConnection(config.Database.Connection)

    if err != nil {
        panic(err)
    }

    rand_asset, err := db_conn.FetchRandomAssetId()
    if err != nil {
        panic(err)
    }

    fmt.Println("Random Asset: ", rand_asset)

    v, err := db_conn.FetchAssetById(9006621)

    if err != nil {
        panic(err)
    }

    // router := mux.NewRouter()
    // router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Asset: %v \n", v)

    m := MakeAssetMessage(v)
    var d []json.RawMessage

    for _, val := range v.Datapoints {
        dm := MakeDatapointMessage(val)
        dmout, err := json.Marshal(dm)

        if err != nil {
            panic(err)
        }

        d = append(d, dmout)
    }

    m.Datapoints = d

    // out, err := json.Marshal(m)

    // if err != nil {
    //     panic(err)
    // }

    sm, err := json.Marshal(m)
    if err != nil {
        panic(err)
    }

    fmt.Println("Message: ", string(sm[:]))
}
