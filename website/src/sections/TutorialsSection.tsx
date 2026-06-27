import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Tag, Typography } from 'antd'
import { ClockCircleOutlined, ReadOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'

const { Title, Paragraph } = Typography

const difficultyColors: Record<string, string> = {
  'Beginner': '#16A34A',
  '入门': '#16A34A',
  'Intermediate': '#2563EB',
  '进阶': '#2563EB',
  'Advanced': '#DC2626',
  '高级': '#DC2626',
}

const TutorialsSection: React.FC = () => {
  const { t } = useTranslation()

  const items = t('tutorials.items', { returnObjects: true }) as Array<{
    category: string
    categoryColor: string
    title: string
    description: string
    difficulty: string
    readTime: string
  }>

  return (
    <section id="tutorials">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('tutorials.title')} subtitle={t('tutorials.subtitle')} />

        <Row gutter={[24, 24]}>
          {items.map((item, index) => (
            <Col xs={24} sm={12} lg={8} key={index}>
              <div className="tutorial-card">
                <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 14 }}>
                  <Tag
                    style={{
                      background: '#F8FAFC',
                      color: item.categoryColor,
                      border: `1px solid ${item.categoryColor}30`,
                      fontWeight: 600,
                      borderRadius: 2,
                      margin: 0,
                      fontSize: 12,
                    }}
                  >
                    {item.category}
                  </Tag>
                  <Tag
                    style={{
                      background: '#F8FAFC',
                      color: difficultyColors[item.difficulty] || '#475569',
                      border: `1px solid ${difficultyColors[item.difficulty] || '#475569'}30`,
                      borderRadius: 2,
                      margin: 0,
                      fontWeight: 500,
                      fontSize: 12,
                    }}
                  >
                    {item.difficulty}
                  </Tag>
                </div>

                <Title level={5} style={{ marginBottom: 8, fontWeight: 700, fontSize: 15, lineHeight: 1.4 }}>
                  <ReadOutlined style={{ marginRight: 6, color: '#2563EB' }} />
                  {item.title}
                </Title>

                <Paragraph style={{ color: '#64748B', marginBottom: 12, lineHeight: 1.6, fontSize: 14 }}>
                  {item.description}
                </Paragraph>

                <div style={{ display: 'flex', alignItems: 'center', color: '#94A3B8', fontSize: 12 }}>
                  <ClockCircleOutlined style={{ marginRight: 4 }} />
                  {item.readTime}
                </div>
              </div>
            </Col>
          ))}
        </Row>
      </div>
    </section>
  )
}

export default TutorialsSection
