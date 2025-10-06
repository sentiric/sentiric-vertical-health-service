# 🏥 Sentiric Vertical Health Service - Mantık ve Akış Mimarisi

**Stratejik Rol:** Sağlık sektörüne özel iş mantığını (doktor arama, randevu planlama, tıbbi sorgulama) sunar. Hasta verilerine (HIPAA/KVKK) uyumlu güvenli bir arayüz sağlar.

---

## 1. Temel Akış: Doktor Arama (FindDoctor)

```mermaid
sequenceDiagram
    participant Agent as Agent Service
    participant VHS as Health Service
    participant EMR as Harici EMR Sistemi (API)
    
    Agent->>VHS: FindDoctor(specialty="Kardiyoloji", location="İstanbul")
    
    Note over VHS: 1. Harici EMR/Sistem Sorgusu
    VHS->>EMR: GET /doctors?specialty=...
    EMR-->>VHS: Doktor Listesi
    
    Note over VHS: 2. Sonuçları Güvenli Mesaja Çevir
    VHS-->>Agent: FindDoctorResponse(doctors: [DoctorInfo, ...])
```

## 2. Hasta Verisi (Confidentiality)

Bu servis, hassas tıbbi verilere erişebileceği için en yüksek düzeyde güvenlik ve yetkilendirme gerektirir. Tüm veritabanı/harici API çağrılarında tenant ve user ID'lerinin yetkilendirme için kullanılması zorunludur.