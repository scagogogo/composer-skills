import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Card, Table, Image, Typography } from 'antd'
import { CloudOutlined, CodeOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'

const { Title } = Typography

const SDKComparisonSection: React.FC = () => {
  const { t } = useTranslation()

  const fields = t('sdkComparison.fields', { returnObjects: true }) as Array<{
    label: string
    packagist: string
    composer: string
  }>

  const columns = [
    {
      title: '',
      dataIndex: 'label',
      key: 'label',
      width: '25%',
      render: (text: string) => <strong>{text}</strong>,
    },
    {
      title: (
        <span>
          <CloudOutlined style={{ marginRight: 8, color: '#0284C7' }} />
          {t('sdkComparison.packagistTitle')}
        </span>
      ),
      dataIndex: 'packagist',
      key: 'packagist',
      width: '37.5%',
    },
    {
      title: (
        <span>
          <CodeOutlined style={{ marginRight: 8, color: '#4F46E5' }} />
          {t('sdkComparison.composerTitle')}
        </span>
      ),
      dataIndex: 'composer',
      key: 'composer',
      width: '37.5%',
    },
  ]

  return (
    <section id="sdk-comparison">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('sdkComparison.title')} subtitle={t('sdkComparison.subtitle')} />

        <div style={{ textAlign: 'center', marginBottom: 40 }}>
          <Image
            src={`${import.meta.env.BASE_URL}images/sdk-comparison.png`}
            alt="SDK Comparison"
            style={{ maxWidth: '90%', borderRadius: 8 }}
            preview={false}
          />
        </div>

        <Row gutter={[24, 24]} style={{ marginBottom: 32 }}>
          <Col xs={24} md={12}>
            <Card
              style={{
                textAlign: 'center',
                background: 'linear-gradient(135deg, #E0F2FE, #BAE6FD)',
                border: 'none',
                borderRadius: 12,
              }}
            >
              <CloudOutlined style={{ fontSize: 36, color: '#0284C7', marginBottom: 12 }} />
              <Title level={4} style={{ color: '#0C4A6E', marginBottom: 8 }}>
                {t('sdkComparison.packagistTitle')}
              </Title>
            </Card>
          </Col>
          <Col xs={24} md={12}>
            <Card
              style={{
                textAlign: 'center',
                background: 'linear-gradient(135deg, #E0E7FF, #C7D2FE)',
                border: 'none',
                borderRadius: 12,
              }}
            >
              <CodeOutlined style={{ fontSize: 36, color: '#4F46E5', marginBottom: 12 }} />
              <Title level={4} style={{ color: '#312E81', marginBottom: 8 }}>
                {t('sdkComparison.composerTitle')}
              </Title>
            </Card>
          </Col>
        </Row>

        <Table
          columns={columns}
          dataSource={fields.map((item, i) => ({ ...item, key: i }))}
          pagination={false}
          bordered
        />
      </div>
    </section>
  )
}

export default SDKComparisonSection
