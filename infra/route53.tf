locals {
  zone_id     = "Z08521423L6AXCPPPYPGE"
  root_domain = "developic.kr"
  static_ip   = "15.165.21.53"
}

resource "aws_route53_record" "api" {
  name = join(".", ["api", local.root_domain])
  type = "A"
  ttl  = 300

  zone_id = local.zone_id
  records = [local.static_ip]
}

resource "aws_route53_record" "acme" {
  name = "_acme-challenge.developic.kr"
  type = "TXT"
  ttl  = 300

  zone_id = local.zone_id
  records = ["Ar-_dmgkZePYIdT2xBOdxDCwKYmlxY8Q3i0lhrCZ4cA"]
}