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

const colorMap = ['#4F46E5', '#E11D48', '#0284C7', '#059669', '#D97706', '#7C3AED']

const UseCasesSection: React.FC = () => {
  const { t } = useTranslation()

  const items = t('useCases.items', { returnObjects: true }) as Array<{
    title: string
    description: string
  }>

  return (
    <section id="use-cases" className="section-alt">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('useCases.title')} subtitle={t('useCases.subtitle')} />

        <Row gutter={[24, 24]}>
          {items.map((item, index) => (
            <Col xs={24} sm={12} lg={8} key={index}>
              <Card
                hoverable
                className="feature-card"
                style={{
                  height: '100%',
                  borderRadius: 16,
                  border: '1px solid #E2E8F0',
                  background: '#fff',
                }}
                styles={{
                  body: { padding: 28 },
                }}
              >
                <div
                  style={{
                    width: 52,
                    height: 52,
                    borderRadius: 14,
                    background: `${colorMap[index]}10`,
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    fontSize: 24,
                    color: colorMap[index],
                    marginBottom: 20,
                  }}
                >
                  {iconMap[index]}
                </div>
                <Title level={5} style={{ marginBottom: 10, fontWeight: 700 }}>
                  {item.title}
                </Title>
                <Paragraph style={{ color: '#64748B', marginBottom: 0, lineHeight: 1.6 }}>
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
