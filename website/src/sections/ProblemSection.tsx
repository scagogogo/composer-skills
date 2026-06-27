import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Table, Tag } from 'antd'
import { CloseCircleOutlined, CheckCircleOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import CodeBlock from '../components/CodeBlock'

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
        <span style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
          <CloseCircleOutlined style={{ color: '#DC2626' }} />
          <span>{text}</span>
        </span>
      ),
    },
    {
      title: t('problem.solutionHeader'),
      dataIndex: 'solution',
      key: 'solution',
      render: (text: string) => (
        <span style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
          <CheckCircleOutlined style={{ color: '#16A34A' }} />
          <span>{text}</span>
        </span>
      ),
    },
  ]

  return (
    <section id="problem" className="section-alt">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('problem.title')} subtitle={t('problem.subtitle')} />

        <Row gutter={[24, 24]} style={{ marginBottom: 48 }}>
          <Col xs={24} md={12}>
            <div style={{ border: '1px solid #FECACA', background: '#FEF2F2', borderRadius: 4, padding: '16px 20px 12px' }}>
              <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 10 }}>
                <Tag color="error" style={{ margin: 0, borderRadius: 2, fontWeight: 600 }}>OLD</Tag>
                <span style={{ color: '#991B1B', fontWeight: 600, fontSize: 13 }}>{t('problem.oldWay')}</span>
              </div>
              <div className="code-wrapper">
                <CodeBlock code={t('problem.oldCode')} />
              </div>
            </div>
          </Col>
          <Col xs={24} md={12}>
            <div style={{ border: '1px solid #BBF7D0', background: '#F0FDF4', borderRadius: 4, padding: '16px 20px 12px' }}>
              <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 10 }}>
                <Tag color="success" style={{ margin: 0, borderRadius: 2, fontWeight: 600 }}>NEW</Tag>
                <span style={{ color: '#166534', fontWeight: 600, fontSize: 13 }}>{t('problem.newWay')}</span>
              </div>
              <div className="code-wrapper">
                <CodeBlock code={t('problem.newCode')} />
              </div>
            </div>
          </Col>
        </Row>

        <Table
          columns={tableColumns}
          dataSource={painPoints.map((item, i) => ({ ...item, key: i }))}
          pagination={false}
          bordered
        />
      </div>
    </section>
  )
}

export default ProblemSection
