resource "packetfabric_point_to_point" "ptp1" {
  provider          = packetfabric
  description       = var.pf_description
  speed             = var.pf_ptp_speed
  media             = var.pf_ptp_media
  subscription_term = var.pf_ptp_subterm
  endpoints {
    pop     = var.pf_ptp_pop1
    zone    = var.pf_ptp_zone1
    autoneg = var.pf_ptp_autoneg
  }
  endpoints {
    pop     = var.pf_ptp_pop2
    zone    = var.pf_ptp_zone2
    autoneg = var.pf_ptp_autoneg
  }
}
