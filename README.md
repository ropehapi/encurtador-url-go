# **Encurtador de URLs com Go**
Desenvolvido em 2023 para fins de estudos, o Encurtador de URLs Go é uma ferramenta de dois endpoints onde você encurta e redireciona a URLs.

## **Instalação**
Para subir o projeto localmente em sua máquina, você só precisará do Go e do MySQL instalados. Com isso feito, siga o passo a passo abaixo:

- Baixe o repositório do projeto através do git clone
- Configure o arquivo `config/database.go` de acordo com as suas credenciais de acesso ao banco
- Rode as duas migrations no seu banco de dados
- Suba a aplicação na porta de sua preferência

## **Endpoints**
### **Encurtador de URL**
- Endpoint: /encurta
- Método: POST
- Form: 
    - url: a url que se deseja encurtar
### **Redirecionador**
- Endpoint: /desencurta/\<code>
- Método: GET
