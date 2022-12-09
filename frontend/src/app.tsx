import { Button, Modal } from 'antd';
import React, { useState } from 'react';

const App: React.FC = () => {
    const [isModalOpen, setIsModalOpen] = useState(false);

    return (
        <div className="flex items-center">
            <div className="mx-auto my-36">
                <Button onClick={() => setIsModalOpen(true)}>Noop? Hmm...</Button>
            </div>
            <Modal open={isModalOpen}
                   onOk={() => setIsModalOpen(false)}
                   onCancel={() => setIsModalOpen(false)}
                   okButtonProps={{className: 'bg-primary hover:bg-tertiary'}}
                   title="A Modal"
                   okText="Okie dokie"
            >
                <Button>Why is there another button ???</Button>
            </Modal>
        </div>
    );
};

export default App;
