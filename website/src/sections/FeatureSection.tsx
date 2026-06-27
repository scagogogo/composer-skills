import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col } from 'antd'
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
import FeatureCard from '../components/FeatureCard'

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

        <Row gutter={[24, 24]}>
          {items.map((item, index) => (
            <Col xs={24} sm={12} lg={8} key={index}>
              <FeatureCard
                icon={iconMap[index]}
                title={item.title}
                description={item.description}
              />
            </Col>
          ))}
        </Row>
      </div>
    </section>
  )
}

export default FeatureSection
