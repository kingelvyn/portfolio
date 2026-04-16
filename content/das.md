# NAS (Network Attached Storage)

---

### Summary
> Expanding my [homelab's](https://portfolio.elvynprise.xyz/projects/homelab) storage capabilities

<div style="text-align:center;">
  <img src="/static/images/homelab.jpg" alt="alt text" style="max-width:60%; height:auto; border-radius:8px; box-shadow:0 4px 12px rgba(0,0,0,0.15);">
</div>

___

### Features:

- ZFS pool configured for RAIDZ1 for case of single hard drive failure
- JBOD attached to Proxmox host to allow for network sharing of storage
- Externally powered separate from Proxmox host server

### Upgrade:
- Upgraded to a dedicated PiNAS set-up
- Separate NAS OS via Open Media Vault (OMV)
- [3D-Printed structure](https://makerworld.com/en/models/1605027-raspberry-pi-5-based-4-bay-nas#profileId-1692368)

### BOM:
| Part | Price |
| -------- | -------- |
| RPi 5 4gb | ~$80 |
| Radxa SATA Hat | ~$80 |
| Noctua NF-A8 Fan | ~$18 |
| PSU (12V6A) | ~$16 |
| Threaded Inserts | ~$10 |
| SATA Extenders | ~$20 |
| Total | ~$220 |


### Skills Used / Developed:

- Hardware integration (attached drives externally using a SAS HBA IT card)
- 3D printed custom enclosure
- Managed power delivery and cabling for reliability and performance
- Mounted and formatted drives for a RAIDZ ZFS pool with a single parity drive
- Pass storage into Proxmox for samba sharing over network
