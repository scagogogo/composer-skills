import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Card, Typography } from 'antd'
import {
  CloudServerOutlined,
  SafetyOutlined,
  CloudSyncOutlined,
  DashboardOutlined,
  ToolOutlined,
  BuildOutlined,
} from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'

const { Title, Paragraph } = Typography

const iconMap = [
  <CloudServerOutlined />,
  <SafetyOutlined />,
  <CloudSyncOutlined />,
  <DashboardOutlined />,
  <ToolOutlined />,
  <BuildOutlined />,
]

const colorMap = ['#2563EB', '#DC2626', '#0891B2', '#16A34A', '#EA580C', '#0F172A']

const UseCasesSection: React.FC = () => {
  const { t } = useTranslation()

  const items = t('useCases.items', { returnObjects: true }) as Array<{
    title: string
    description: string
  }>

  return (
    <section id="use-cases">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('useCases.title')} subtitle={t('useCases.subtitle')} />

        <Row gutter={[24, 24]}>
          {items.map((item, index) => (
            <Col xs={24} sm={12} lg={8} key={index}>
              <Card
                className="feature-card"
                style={{
                  height: '100%',
                  background: '#fff',
                }}
                styles={{
                  body: { padding: 20 },
                }}
              >
                <div
                  style={{
                    width: 36,
                    height: 36,
                    borderRadius: 4,
                    background: `${colorMap[index]}0A`,
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    fontSize: 18,
                    color: colorMap[index],
                    marginBottom: 14,
                  }}
                >
                  {iconMap[index]}
                </div>
                <Title level={5} style={{ marginBottom: 8, fontWeight: 700, fontSize: 15 }}>
                  {item.title}
                </Title>
                <Paragraph style={{ color: '#64748B', marginBottom: 0, lineHeight: 1.6, fontSize: 14 }}>
                  {item.description}
                </Paragraph>
              </Card>
            </Col>
          ))}
        </Row>
      </div>
    </section>
  )
}

export default UseCasesSection
