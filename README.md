# KilatMesh Agent

**KilatMesh Agent** adalah agen konektivitas sekolah dari
**PT Kilat Cloud Indonesia**. Agen ini berjalan sebagai service di PC
sekolah (Windows/Linux) dan menghubungkan jaringan lokal sekolah —
terutama PC Dapodik — ke platform EDU Kilat Cloud melalui mesh VPN
berbasis WireGuard, tanpa membuka port inbound apa pun di sekolah.

Dokumen desain integrasi: `docs/NETBIRD-INTEGRATION.md` di repo
platform `EDU.KILAT-CLOUD`.

## Cara kerja singkat

1. Operator platform membuat *setup key* khusus untuk sekolah Anda
   dari server management KilatMesh (satu key per sekolah).
2. Install agen ini di PC Dapodik (atau PC khusus di LAN sekolah).
3. Daftarkan PC ke mesh dengan satu perintah:

   ```bash
   # Linux (sebagai root)
   kilatmesh up --management-url https://mesh.<domain-platform> \
                --setup-key <KEY-SEKOLAH-ANDA> \
                --hostname agent-<kode-sekolah>

   # Windows (PowerShell sebagai Administrator, perintah yang sama)
   kilatmesh up --management-url https://mesh.<domain-platform> `
                --setup-key <KEY-SEKOLAH-ANDA> `
                --hostname agent-<kode-sekolah>
   ```

4. Selesai. Platform kemudian dapat menjangkau WebService Dapodik di
   sekolah Anda lewat tunnel terenkripsi, sesuai policy akses
   per-sekolah yang diatur operator (default: ditolak semua, dibuka
   per kebutuhan).

## Service

| | Linux | Windows |
|---|---|---|
| Nama service | `kilatmesh` (systemd) | `kilatmesh` |
| Nama tampil | KilatMesh Agent | KilatMesh Agent |
| Kelola | `kilatmesh service install|start|stop|status` | sama |

Path penting (Linux): konfigurasi `/etc/kilatmesh/` dan
`/var/lib/kilatmesh/`, log `/var/log/kilatmesh/client.log`,
socket daemon `/var/run/kilatmesh.sock`.
Di Windows: `%PROGRAMDATA%\KilatMesh\`.

## Build dari source

```bash
source ~/.dev-env.sh   # toolchain Go di lingkungan Kilat Cloud

# Linux amd64
cd client && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -tags load_wgnt_from_rsrc \
  -ldflags "-s -w -X github.com/netbirdio/netbird/version.version=0.82.1" \
  -o ../dist/kilatmesh-linux-amd64 .

# Windows amd64
cd client && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build \
  -tags load_wgnt_from_rsrc \
  -ldflags "-s -w -X github.com/netbirdio/netbird/version.version=0.82.1" \
  -o ../dist/kilatmesh.exe .
```

## Asal-usul & lisensi

KilatMesh Agent adalah **fork dari NetBird** (client, lisensi
BSD-3-Clause), dipin pada rilis upstream **v0.82.1** — lihat
[UPSTREAM.md](UPSTREAM.md) untuk detail pin dan ringkasan lisensi
per komponen, [NOTICE](NOTICE) untuk atribusi fork, dan
[README.upstream.md](README.upstream.md) untuk README upstream asli.

Yang diganti dari upstream hanyalah branding yang terlihat pengguna
(nama produk, string CLI, nama service, path default, nama objek
firewall, dokumentasi). Protokol, format key/identitas, dan
kompatibilitas dengan server management NetBird standar **tidak
diubah** — agen ini tetap klien NetBird yang berfungsi penuh.

Komponen server NetBird (management/signal/relay/dashboard)
berlisensi AGPLv3, **bukan** bagian dari fork ini, dan dijalankan
apa adanya (self-hosted, tanpa modifikasi). KilatMesh bukan produk
resmi NetBird dan tidak didukung/diendors oleh NetBird GmbH.
