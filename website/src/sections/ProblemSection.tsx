import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Table, Typography } from 'antd'
import { CloseCircleOutlined, CheckCircleOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import CodeBlock from '../components/CodeBlock'

const { Title } = Typography

const ProblemSection: React.FC = () => {
  const { t } = useTranslation()

  const painPoints = t('problem.painPoints', { returnObjects: true }) as Array<{
    pain: string
    solution: string
  }>

  const tableColumns = [
    {
      title: t('problem.painHeader'),
      dataIndex: 'pain',
      key: 'pain',
      render: (text: string) => (
        <span>
          <CloseCircleOutlined style={{ color: '#E11D48', marginRight: 8 }} />
          {text}
        </span>
      ),
    },
    {
      title: t('problem.solutionHeader'),
      dataIndex: 'solution',
      key: 'solution',
      render: (text: string) => (
        <span>
          <CheckCircleOutlined style={{ color: '#059669', marginRight: 8 }} />
          {text}
        </span>
      ),
    },
  ]

  return (
    <section id="problem" style={{ background: '#F8FAFC' }}>
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('problem.title')} subtitle={t('problem.subtitle')} />

        <Row gutter={[32, 32]} style={{ marginBottom: 48 }}>
          <Col xs={24} md={12}>
            <Title level={5} style={{ color: '#E11D48', marginBottom: 12 }}>
              😩 {t('problem.oldWay')}
            </Title>
            <CodeBlock code={t('problem.oldCode')} />
          </Col>
          <Col xs={24} md={12}>
            <Title level={5} style={{ color: '#059669', marginBottom: 12 }}>
              😊 {t('problem.newWay')}
            </Title>
            <CodeBlock code={t('problem.newCode')} />
          </Col>
        </Row>

        <Table
          columns={tableColumns}
          dataSource={painPoints.map((item, i) => ({ ...item, key: i }))}
          pagination={false}
          bordered
          style={{ borderRadius: 8 }}
        />
      </div>
    </section>
  )
}

export default ProblemSection
