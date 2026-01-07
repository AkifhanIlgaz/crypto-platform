### Adım 1: Projeyi İndirin

```bash
git clone https://github.com/AkifhanIlgaz/crypto-platform.git
cd crypto-platform
```

### Adım 2: ENV Dosyasını Oluşturun ve Gerekli Bilgiler Girin

```bash
cp .env.example .env
```

### Adım 3: Backendi Çalıştırın

```bash
docker compose up -d
```

### Adım 4: Frontendi Çalıştırın

```bash
cd frontend
npm install
npm run dev
```

# http://localhost:3000 adresinden erişebilirsiniz
