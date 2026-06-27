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
          <CloudOutlined style={{ marginRight: 6, color: '#0891B2' }} />
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
          <CodeOutlined style={{ marginRight: 6, color: '#2563EB' }} />
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

        <div style={{ textAlign: 'center', marginBottom: 32 }}>
          <Image
            src={`${import.meta.env.BASE_URL}images/sdk-comparison.png`}
            alt="SDK Comparison"
            style={{ maxWidth: '90%' }}
            preview={false}
          />
        </div>

        <Row gutter={[24, 16]} style={{ marginBottom: 24 }}>
          <Col xs={24} md={12}>
            <Card
              style={{
                textAlign: 'center',
                background: '#F0F9FF',
                border: '1px solid #BAE6FD',
                borderRadius: 4,
              }}
              styles={{ body: { padding: 20 } }}
            >
              <CloudOutlined style={{ fontSize: 28, color: '#0891B2', marginBottom: 8 }} />
              <Title level={5} style={{ color: '#0C4A6E', marginBottom: 0, fontSize: 15 }}>
                {t('sdkComparison.packagistTitle')}
              </Title>
            </Card>
          </Col>
          <Col xs={24} md={12}>
            <Card
              style={{
                textAlign: 'center',
                background: '#EFF6FF',
                border: '1px solid #BFDBFE',
                borderRadius: 4,
              }}
              styles={{ body: { padding: 20 } }}
            >
              <CodeOutlined style={{ fontSize: 28, color: '#2563EB', marginBottom: 8 }} />
              <Title level={5} style={{ color: '#1E3A5F', marginBottom: 0, fontSize: 15 }}>
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
