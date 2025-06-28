terraform {
  required_providers {
    ec2instancetype = {
      source  = "local/custom/ec2-instance-type"
      version = "1.0.1"
    }
  }
}

provider "ec2instancetype" {
  access_key = "ASIAVJR3DEEEVCKO3QCD"
  secret_key = "vJFaBFW7ZhGWFI7IR/e+7YPP7KwPn38MT3QrKi7y"
  session_token = "IQoJb3JpZ2luX2VjEJX//////////wEaCWV1LXdlc3QtMSJHMEUCIQD4uQAq+I2j7yRallFAVQCr37UMNVlq3YXxfbTY+6QsEwIgZSiAkVfILryLoipNqvOXO2D1sgVAySjVy9X5U4xjE0IqnAMIjv//////////ARADGgwzNjQxMjI0MTUzNjkiDN1WTlnd37eG83OZkirwAl3Cko6SED2fRoUKiMLUqV/7F/SHfVCJOWhSLT3tlBR7yWbZz4yua6muR3p3Y0RiTKWxi/37/5WkaYZv7WfhLPvXfFRVy37aoFpibsJMPst8JoqK8gXW602QgVvlrCqDh0RiIuUO3rQtgteL96G4pPo/PzCfA1ajwGfencfqA08O9heKpmvcanvm0cycO4wro6hUIpY2ZUOiLPG+B5DzcW/QjFR+mn9485/RzIU/ClODe2MVRD7vNn2jtG6r7Lm39YFOoOHTupap1cw5FSfijc2QOcwJ0sRAmhUzmcWCDqrHru1rjXAdgTmapwjRG1FiXeEkKkv6Ehz0lqDJcxQw0e/uDF5UdBP56BON/7IBlukVGBfxpbGS+wRKSEsVPaPhd1YzX3i5TSgAgWb4UhnwcXk7A3htQOm42pZbhORFFqCKcKalDkoj1HpGXqmyh/WmlYmVi4xTEvUXU74JPQ6foZkizLQi08NcodEYLCp7WjAzMNvK/8IGOqYBYMCbaiXXEtd0ipT8tzWQSQwxXvDaqirjoVGrTDFct0Js2DviHoCgpZ4j1X8s3mhn9OijWtTss2yFzZ56XYRc9SASUU2uJbRHJe/NGRqCIPZ6qGIn1AsJXfdcSY+KtA3MGgGl9Y9M0cGu3aqJPPGbZg+uPJROkdMqwnwAtTwgbr91Cx9W+4Rt949AOhteGV4vTZPlXItUQ8wDO7iWBYF0WRigM78E/w=="
  region = "eu-west-1"
}