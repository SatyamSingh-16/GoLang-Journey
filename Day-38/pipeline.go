package day38
pipe := rdb.Pipeline()

setCmd := pipe.Set(ctx, "name", "Satyam", 0)
getCmd := pipe.Get(ctx, "name")

_, err := pipe.Exec(ctx)
if err != nil {
    return err
}

fmt.Println(setCmd.Val())
fmt.Println(getCmd.Val())