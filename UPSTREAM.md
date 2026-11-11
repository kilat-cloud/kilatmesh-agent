# UPSTREAM — Sumber Fork KilatMesh Agent

Dokumen ini mencatat asal-usul kode di repositori ini agar sinkronisasi
dengan upstream selalu bisa dilakukan secara sadar dan terlacak.

## Sumber

| Item | Nilai |
|---|---|
| Repositori upstream | https://github.com/netbirdio/netbird |
| Remote lokal | `upstream` → `https://github.com/netbirdio/netbird` (fetch/push) |
| Tag yang dipin | **v0.82.1** (rilis stabil terakhir saat fork dibuat) |
| Commit tag | `a48050004588a5651a4191bfa4e363da5ee273a3` |
| Tanggal commit upstream | 2026-11-10 |
| Tanggal fork dibuat | 2026-11-11 |
| Remote `origin` | **Tidak diset.** Fork ini tidak dipush ke mana pun untuk saat ini. Bila nanti dipublish, `origin` akan menunjuk ke repositori milik PT Kilat Cloud Indonesia, bukan ke `netbirdio`. |

## Ringkasan lisensi per komponen (sesuai file LICENSE upstream)

- **BSD-3-Clause**: seluruh repositori, **kecuali** direktori di bawah.
  Ini mencakup `client/` — satu-satunya komponen yang di-rebrand dan
  didistribusikan oleh fork ini sebagai **KilatMesh Agent**.
- **AGPLv3**: direktori `management/`, `signal/`, `relay/`, dan
  `combined/`. Komponen-komponen ini **tidak** dimodifikasi oleh fork
  ini; untuk go-live, server upstream dijalankan apa adanya
  (self-hosted, tanpa modifikasi). Setiap fork atas komponen AGPL di
  masa depan **wajib** berada di repositori publik sesuai ketentuan
  AGPLv3.

File `LICENSE`, direktori `LICENSES/`, dan `AUTHORS` upstream
dipertahankan apa adanya. Lihat juga `NOTICE` untuk atribusi fork.

## Kebijakan sinkronisasi

- Tarik perubahan upstream hanya dari remote `upstream`
  (`git fetch upstream`), lalu rebase/merge secara eksplisit per rilis.
- Perubahan rebrand KilatMesh dijaga tetap dangkal dan terlokalisasi
  (string CLI, nama service, path default, nama tabel/chain firewall,
  README/NOTICE) agar rebase ke rilis upstream berikutnya murah.
  Modul Go dan path impor (`github.com/netbirdio/netbird/...`)
  **sengaja tidak diganti nama** — itu internal dan menggantinya akan
  merusak build serta menyulitkan sinkronisasi.
- Protokol wire, format key/identitas peer, dan mekanisme URL
  management **tidak diubah**: KilatMesh Agent tetap klien yang
  sepenuhnya kompatibel dengan server management NetBird standar.
