# okteto-example

sample application: [github.com/okteto/movies](https://github.com/okteto/movies)

```
okteto deploy
```


## development

開発用コンテナを起動

```bash
okteto up
```

開発用コンテナを破棄

```bash
okteto down
```

開発用コンテナ起動後は、新たなコンテナが起動され、ファイルがマウントされているため、ファイルの変更が即時適用される。\
コンテナは起動しており残り続けるが、内部のアプリケーションは起動していないので、手動で起動することで外部からアクセス可能になる。
