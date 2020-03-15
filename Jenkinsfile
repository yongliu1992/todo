// todo

String image    = env.JOB_NAME.split('/')[1].toLowerCase()
String registry = "yongliu1992/$image"
String tag      = "latest"

pipeline {
    agent {
    kubernetes {
      	cloud 'kubernetes'
      	label 'todo'
      	defaultContainer 'jnlp'
      	yamlFile 'pod.yaml'
      }
    }

    stages {
        stage("publish") {
            steps {
                container(name: 'yongliu1992', shell: '/busybox/sh') {
                     withEnv(['PATH+EXTRA=/busybox:/kaniko']) {
                      sh """#!/busybox/sh
                        /kaniko/executor --context=$workspace --destination $registry:$tag
                      """
                     }
                }
            }
        }
    }
}
