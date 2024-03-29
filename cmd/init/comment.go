package init

const ruleComment = `
rules:
  - name: ebs_volume_are_unencrypted
    query: |
      SELECT 
        *
      FROM 
        aws_ec2_ebs_volumes 
      WHERE 
        encrypted = FALSE
    labels:
      tag:
        - Security
        - Misconfigure
      author: Selefra
      standard: Custom
    metadata:
      id: SF010302
      severity: Low
      provider: AWS
      resource_type: EC2
      resource_account_id : '{{.account_id}}'
      resource_id: '{{.id}}'
      resource_region: '{{.availability_zone}}'
      remediation: remediation/ebs_volume_are_unencrypted.md
      title: EBS volume are unencrypted
      description: Ensure that EBS volumes are encrypted.
    output: 'EBS volume are unencrypted, EBS id: {{.id}}, availability zone: {{.availability_zone}}'
`

const moduleComment = `
modules:
  - name: AWS_Security_Demo
    uses:
    - ./rules/iam_mfa.yaml
`
