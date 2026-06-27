import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Card, Image, Typography } from 'antd'
import { SafetyOutlined, BugOutlined, FileProtectOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import CodeBlock from '../components/CodeBlock'

const { Title } = Typography

const SecuritySection: React.FC = () => {
  const { t } = useTranslation()

  const cards = [
    {
      icon: <BugOutlined />,
      title: t('security.auditTitle'),
      code: t('security.auditCode'),
    },
    {
      icon: <SafetyOutlined />,
      title: t('security.remoteTitle'),
      code: t('security.remoteCode'),
    },
    {
      icon: <FileProtectOutlined />,
      title: t('security.validateTitle'),
      code: t('security.validateCode'),
    },
  ]

  return (
    <section id="security" style={{ background: '#F8FAFC' }}>
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('security.title')} subtitle={t('security.subtitle')} />

        <div style={{ textAlign: 'center', marginBottom: 40 }}>
          <Image
            src={`${import.meta.env.BASE_URL}images/security-features.png`}
            alt="Security Features"
            style={{ maxWidth: '90%', borderRadius: 8 }}
            preview={false}
          />
        </div>

        <Row gutter={[24, 24]}>
          {cards.map((card, index) => (
            <Col xs={24} md={8} key={index}>
              <Card
                style={{
                  height: '100%',
                  borderRadius: 12,
                  border: '1px solid #E2E8F0',
                }}
                styles={{
                  body: { padding: 24 },
                }}
              >
                <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 16 }}>
                  <span style={{ fontSize: 20, color: '#E11D48' }}>{card.icon}</span>
                  <Title level={5} style={{ margin: 0 }}>
                    {card.title}
                  </Title>
                </div>
                <CodeBlock code={card.code} />
              </Card>
            </Col>
          ))}
        </Row>
      </div>
    </section>
  )
}

export default SecuritySection
