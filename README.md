# exercism

# Para lidar com a conexão no github.

1. crie a chave publica e a privada com: ssh-keygen -t rsa -b 4096 -C "email@cadastrado.no.git"
2. copie a chave publica para o github e cadastre no settings
3. crie um arquivo de nome config no ~/.ssh e coloque:

#cirobispo@zoho.com account for github
Host github.com-cirobispo
    HostName github.com
    User git
    IdentityFile ~/.ssh/cirobispo_rsa
    IdentitiesOnly yes

4. inicie o ssh-agent com: exec ssh-agent bash
5. coloque a chave privada na memoria do agent com: ssh-add ~/.ssh/arquivo_chave_privada
6. teste se a chave está adicionada com: ssh -Tv github.com-cirobispo
7. crie, no diretorio ~o arquivo .gitconfig e coloque:

[filter "ifs"]
[includeIf "gitdir:diretorio/do/projeto/usuarioA/"]
    path=./.gitconfig-A

[includeIf "gitdir:diretorio/do/projeto/usuarioB/"]
    path=./.gitconfig-B

[user]
    name = Ciro Bispo
    email = cirobispo@zoho.com

[core]
    sshCommand = ssh.exe

8. crie, no mesmo local o arquivo gitconfig-A e o gitconfig-B, etc e coloque:
[user]
    name = Ciro
    email = cirobispo@zoho.com

[core]
    sshCommand = "ssh -i ~/.ssh/arquivo da chave privada correspondente