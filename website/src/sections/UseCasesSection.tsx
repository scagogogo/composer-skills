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

const UseCasesSection: React.FC = () => {
  const { t } = useTranslation()

  const items = t('useCases.items', { returnObjects: true }) as Array<{
    title: string
    description: string
  }>

  return (
    <section id="use-cases" style={{ background: '#F8FAFC' }}>
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('useCases.title')} subtitle={t('useCases.subtitle')} />

        <Row gutter={[24, 24]}>
          {items.map((item, index) => (
            <Col xs={24} sm={12} lg={8} key={index}>
              <Card
                hoverable
                style={{
                  height: '100%',
                  borderRadius: 12,
                  border: '1px solid #E2E8F0',
                }}
                styles={{
                  body: { padding: 24 },
                }}
              >
                <div
                  style={{
                    width: 48,
                    height: 48,
                    borderRadius: 12,
                    background: 'linear-gradient(135deg, #EEF2FF, #E0E7FF)',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    fontSize: 22,
                    color: '#4F46E5',
                    marginBottom: 16,
                  }}
                >
                  {iconMap[index]}
                </div>
                <Title level={5} style={{ marginBottom: 8 }}>
                  {item.title}
                </Title>
                <Paragraph style={{ color: '#475569', marginBottom: 0 }}>
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
