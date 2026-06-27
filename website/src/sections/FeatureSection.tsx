import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Typography } from 'antd'
import {
  CodeOutlined,
  ApiOutlined,
  SafetyOutlined,
  DownloadOutlined,
  LaptopOutlined,
  CodeSandboxOutlined,
  DatabaseOutlined,
  ThunderboltOutlined,
  BookOutlined,
  CheckCircleOutlined,
} from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import { StaggerGrid, StaggerItem } from '../components/ScrollReveal'

const { Title, Paragraph } = Typography

const iconMap = [
  <CodeOutlined />,
  <ApiOutlined />,
  <SafetyOutlined />,
  <DownloadOutlined />,
  <LaptopOutlined />,
  <CodeSandboxOutlined />,
  <DatabaseOutlined />,
  <ThunderboltOutlined />,
  <BookOutlined />,
  <CheckCircleOutlined />,
]

const accentColors = [
  '#2563EB', '#0891B2', '#DC2626', '#16A34A', '#EA580C',
  '#0F172A', '#0284C7', '#D97706', '#059669', '#475569',
]

const FeatureSection: React.FC = () => {
  const { t } = useTranslation()

  const items = t('features.items', { returnObjects: true }) as Array<{
    title: string
    description: string
  }>

  return (
    <section id="features">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('features.title')} subtitle={t('features.subtitle')} />

        <StaggerGrid>
          <Row gutter={[20, 20]}>
            {items.map((item, index) => {
              const color = accentColors[index]
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
                      <div
                        style={{
                          width: 40,
                          height: 40,
                          borderRadius: 4,
                          background: `${color}0A`,
                          display: 'flex',
                          alignItems: 'center',
                          justifyContent: 'center',
                          fontSize: 20,
                          color: color,
                          marginBottom: 16,
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

export default FeatureSection
