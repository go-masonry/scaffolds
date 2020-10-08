# GRPC-Service template

[Cookie Cutter](https://cookiecutter.readthedocs.io) template to create a workshop service from the [Tutorial](https://github.com/go-masonry/tutorial)

Clone scaffolds repository to your local directory

```shell
git clone https://github.com/go-masonry/scaffolds.git
```

Run if you want to create a new directory

```shell
cookiecutter -o $GOPATH/src/github.com/your-organization local-scaffolds-repo/tutorial-example
```

Run if you have an existing directory:

```shell
cookiecutter -f -o $GOPATH/src/github.com/your-organization local-scaffolds-repo/tutorial-example
```