```
<UserList>
    <User>
        <Identifier>İlgili Firmanın Vergi Numarası</Identifier>
        <Alias>e-Fatura mail adresi</Alias>
        <Title>İlgili firmanın ünvanı</Title>
        <Type>E-fatura Tipi</Type>
        <FirstCreationTime>E-Faturaya dahil olduğu ilk tarih</FirstCreationTime>
    </User>
</UserList>
```



**Users.xml Dosyasının Tanımı**
    Firmalar satış faturası kesecekleri zaman, eğer fatura kesilecek olan firma e-faturaya dahil ise ilgili firmaya kağıt fatura yerine e-fatura göndermek durumundadır. Bunu anlamak için devlet users.xml dosyasını hergün güncelleyerek kendi sistemine ekler. Firmaların e-fatura sistemleri bu dosyayı alarak ellerindeki müşteri veya satısı kayıtlarından vergi numaralarını bu dosyada ararlar. Eğer bulurlarsa ilgili müşteri veya satıcıyı e-faturaya dahil olarak işaretlerler ve e-fatura mail adresini ve e-fatura tipi alanlarını güncellerler.

1. Vergi Numarasından kaydı bulma servisi yapılacak.

            `http://localhost:1455/api/getUserFromTaxRegistrationNo/:id`


        eğer kayıt bulamaz ise httpstatus:404 pagenotfound

        eğer kayıt bulunursa httpstatus:200 OK ve User Struct tipinde datayı geri döndürecek.

