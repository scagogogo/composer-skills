import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Typography } from 'antd'
import {
  CloudServerOutlined,
  SafetyOutlined,
  CloudSyncOutlined,
  DashboardOutlined,
  ToolOutlined,
  BuildOutlined,
} from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import { StaggerGrid, StaggerItem } from '../components/ScrollReveal'

const { Paragraph } = Typography

const iconMap = [
  <CloudServerOutlined />,
  <SafetyOutlined />,
  <CloudSyncOutlined />,
  <DashboardOutlined />,
  <ToolOutlined />,
  <BuildOutlined />,
]

const colorMap = ['#2563EB', '#DC2626', '#0891B2', '#16A34A', '#EA580C', '#475569']

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

        <StaggerGrid>
          <Row gutter={[20, 20]}>
            {items.map((item, index) => {
              const color = colorMap[index]
              return (
                <Col xs={24} sm={12} lg={8} key={index}>
                  <StaggerItem>
                    <div
                      className="feature-card"
                      style={{
                        height: '100%',
                        background: '#fff',
                        border: '1px solid #E2E8F0',
                        borderRadius: 4,
                        padding: 24,
                        transition: 'border-color 0.2s ease, transform 0.2s ease',
                      }}
                      onMouseEnter={(e) => {
                        e.currentTarget.style.borderColor = color
                        e.currentTarget.style.transform = 'translateY(-2px)'
                      }}
                      onMouseLeave={(e) => {
                        e.currentTarget.style.borderColor = '#E2E8F0'
                        e.currentTarget.style.transform = 'translateY(0)'
                      }}
                    >
                      {/* Icon with colored left border accent */}
                      <div style={{ display: 'flex', gap: 14, marginBottom: 14 }}>
                        <div
                          style={{
                            width: 4,
                            borderRadius: 2,
                            background: color,
                            flexShrink: 0,
                            minHeight: 36,
                          }}
                        />
                        <div
                          style={{
                            width: 36,
                            height: 36,
                            borderRadius: 4,
                            background: `${color}0A`,
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center',
                            fontSize: 18,
                            color: color,
                          }}
                        >
                          {iconMap[index]}
                        </div>
                      </div>
                      <div style={{ fontWeight: 700, fontSize: 15, color: '#0F172A', marginBottom: 8 }}>
                        {item.title}
                      </div>
                      <Paragraph style={{ color: '#64748B', marginBottom: 0, lineHeight: 1.6, fontSize: 14 }}>
                        {item.description}
                      </Paragraph>
                    </div>
                  </StaggerItem>
                </Col>
              )
            })}
          </Row>
        </StaggerGrid>
      </div>
    </section>
  )
}

export default UseCasesSection
