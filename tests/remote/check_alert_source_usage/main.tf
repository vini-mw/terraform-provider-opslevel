resource "opslevel_check_alert_source_usage" "test" {
  alert_name_predicate = {
    type  = var.alert_name_predicate.type
    value = var.alert_name_predicate.value
  }
  alert_type = var.alert_type

  # -- check base fields --
  category  = var.category
  enable_on = var.enable_on
  enabled   = var.enabled
  filter    = var.filter
  level     = var.level
  name      = var.name
  notes     = var.notes
  owner     = var.owner
}
