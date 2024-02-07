clear
echo "Gerando branches para todos os subdiretorios"

echo "git status | grep -i '/$' > result.txt"

for item in $(cat result.txt); do 
    echo "Criando branch $item e deixando ele ativo"
    echo
    git branch "$item"
    git checkout "$item"
    git branch
    echo "Adicionando arquivos no $item"
    echo
    git add "$item"
    git commit -m "$item"
    echo "Voltando para a branch main"
    echo
    git checkout main
    git rebase "$item"
    git branch
done
