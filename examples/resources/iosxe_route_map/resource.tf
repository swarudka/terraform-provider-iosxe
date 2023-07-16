resource "iosxe_route_map" "example" {
  name = "RM1"
  entries = [
    {
      seq                                      = 10
      operation                                = "permit"
      description                              = "Entry 10"
      continue                                 = false
      match_interfaces                         = ["Loopback1"]
      match_ip_address_access_lists            = ["ACL1"]
      match_ip_next_hop_access_lists           = ["ACL1"]
      match_ipv6_address_access_lists          = "ACL1"
      match_ipv6_next_hop_access_lists         = "ACL1"
      match_route_type_external                = true
      match_route_type_external_type_1         = true
      match_route_type_external_type_2         = true
      match_route_type_internal                = true
      match_route_type_level_1                 = true
      match_route_type_level_2                 = true
      match_route_type_local                   = true
      match_source_protocol_bgp                = ["65000"]
      match_source_protocol_connected          = true
      match_source_protocol_eigrp              = ["10"]
      match_source_protocol_isis               = true
      match_source_protocol_lisp               = true
      match_source_protocol_ospf               = ["10"]
      match_source_protocol_ospfv3             = ["10"]
      match_source_protocol_rip                = true
      match_source_protocol_static             = true
      match_tags                               = [100]
      match_track                              = 1
      match_as_paths_legacy                    = [10]
      match_community_lists_legacy             = ["COMM1"]
      match_extcommunity_lists_legacy          = ["EXTCOMM1"]
      match_local_preferences_legacy           = [100]
      set_default_interfaces                   = ["Loopback1"]
      set_global                               = false
      set_interfaces                           = ["Loopback1"]
      set_ip_address                           = "PFL1"
      set_ip_default_global_next_hop_address   = ["1.2.3.4"]
      set_ip_default_next_hop_address          = ["1.2.3.4"]
      set_ip_global_next_hop_address           = ["1.2.3.4"]
      set_ip_next_hop_address                  = ["1.2.3.4"]
      set_ip_qos_group                         = 1
      set_ipv6_address                         = ["PFL2"]
      set_ipv6_default_global_next_hop         = "2001::1"
      set_ipv6_default_next_hop                = ["2001::1"]
      set_ipv6_next_hop                        = ["2001::1"]
      set_level_1                              = true
      set_metric_value                         = 110
      set_metric_delay                         = "10"
      set_metric_reliability                   = 90
      set_metric_loading                       = 10
      set_metric_mtu                           = 1500
      set_metric_type                          = "external"
      set_tag                                  = 100
      set_as_path_prepend_as_legacy            = "65001 65001"
      set_as_path_prepend_last_as_legacy       = 5
      set_as_path_tag_legacy                   = true
      set_communities_legacy                   = ["1:2"]
      set_communities_additive_legacy          = true
      set_community_list_delete_legacy         = true
      set_community_list_name_legacy           = "COMML1"
      set_extcomunity_rt_legacy                = ["10:10"]
      set_extcomunity_soo_legacy               = "10:10"
      set_extcomunity_vpn_distinguisher_legacy = "10:10"
      set_local_preference_legacy              = 110
      set_weight_legacy                        = 10000
    }
  ]
}