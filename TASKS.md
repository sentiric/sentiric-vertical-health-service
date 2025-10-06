# 🏥 Sentiric Vertical Health Service - Görev Listesi

Bu servisin mevcut ve gelecekteki tüm geliştirme görevleri, platformun merkezi görev yönetimi reposu olan **`sentiric-tasks`**'ta yönetilmektedir.

➡️ **[Aktif Görev Panosuna Git](https://github.com/sentiric/sentiric-tasks/blob/main/TASKS.md)**

---
Bu belge, servise özel, çok küçük ve acil görevler için geçici bir not defteri olarak kullanılabilir.

## Faz 1: Minimal İşlevsellik (INFRA-02)
- [x] Temel Go projesi ve Dockerfile oluşturuldu.
- [x] gRPC sunucusu iskeleti (`FindDoctor`, `ScheduleAppointment`) eklendi.
- [ ] Sahte EMR (Elektronik Tıbbi Kayıt) adaptörü implemente edilecek. (VBS-HEALTH-01)
- [ ] Hasta Kimliği doğrulama mantığı (user-service üzerinden) entegre edilecek. (VBS-HEALTH-02)