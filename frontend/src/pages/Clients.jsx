import React, { useState, useEffect } from 'react';
import { Plus, Trash2, Settings, AlertCircle, Info, Code, Save, X } from 'lucide-react';
import { clientsAPI, mappingAPI } from '../services/api';
import toast from 'react-hot-toast';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Badge } from '@/components/ui/badge';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Textarea } from '@/components/ui/textarea';

const Clients = () => {
  const [clients, setClients] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showCreateModal, setShowCreateModal] = useState(false);
  const [newClientName, setNewClientName] = useState('');
  const [creating, setCreating] = useState(false);
  const [selectedClient, setSelectedClient] = useState(null);
  const [mappings, setMappings] = useState([]);
  const [showMappingForm, setShowMappingForm] = useState(false);
  const [newMapping, setNewMapping] = useState({
    source_path: '',
    destination_path: '',
    transform_type: 'copy',
    transform_logic: '',
    default_value: '',
    required: false
  });
  const [savingMapping, setSavingMapping] = useState(false);

  useEffect(() => {
    loadClients();
  }, []);

  const loadClients = async () => {
    try {
      const data = await clientsAPI.getAll();
      setClients(data);
    } catch (error) {
      toast.error('Failed to load clients');
    } finally {
      setLoading(false);
    }
  };

  const handleCreateClient = async (e) => {
    e.preventDefault();
    if (!newClientName.trim()) return;

    setCreating(true);
    try {
      await clientsAPI.create(newClientName.trim());
      toast.success('Client created successfully');
      setNewClientName('');
      setShowCreateModal(false);
      loadClients();
    } catch (error) {
      toast.error('Failed to create client');
    } finally {
      setCreating(false);
    }
  };

  const handleDeleteClient = async (clientId, clientName) => {
    if (!confirm(`Are you sure you want to delete "${clientName}"? This will also delete all associated mapping rules.`)) {
      return;
    }

    try {
      await clientsAPI.delete(clientId);
      toast.success('Client deleted successfully');
      loadClients();
    } catch (error) {
      toast.error('Failed to delete client');
    }
  };

  const loadMappings = async (clientId) => {
    try {
      const data = await mappingAPI.getByClient(clientId);
      setMappings(data);
    } catch (error) {
      toast.error('Failed to load mapping rules');
      setMappings([]);
    }
  };

  const handleViewMappings = (client) => {
    setSelectedClient(client);
    loadMappings(client.id);
  };

  const handleCreateMapping = async (e) => {
    e.preventDefault();
    if (!selectedClient) return;

    setSavingMapping(true);
    try {
      // Convert string paths to arrays
      const mappingRule = {
        ...newMapping,
        source_path: newMapping.source_path.split('.').filter(p => p.trim()),
        destination_path: newMapping.destination_path.split('.').filter(p => p.trim())
      };

      await mappingAPI.create(selectedClient.id, [mappingRule]);
      toast.success('Mapping rule created successfully');
      setShowMappingForm(false);
      setNewMapping({
        source_path: '',
        destination_path: '',
        transform_type: 'copy',
        transform_logic: '',
        default_value: '',
        required: false
      });
      loadMappings(selectedClient.id);
    } catch (error) {
      toast.error('Failed to create mapping rule');
    } finally {
      setSavingMapping(false);
    }
  };

  const openExpressionHelp = () => {
    window.open('/expression-help', '_blank');
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600"></div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between border-b border-gray-200 pb-5">
        <div>
          <h1 className="text-2xl font-bold leading-7 text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">
            Clients
          </h1>
          <p className="mt-1 text-sm text-gray-500">
            Manage API clients and their mapping rules
          </p>
        </div>
        <Button
          onClick={() => setShowCreateModal(true)}
          className="flex items-center"
        >
          <Plus className="h-5 w-5 mr-2" />
          New Client
        </Button>
      </div>

      {/* Clients Table */}
      <Card>
        <CardContent className="p-0">
          {clients.length === 0 ? (
            <div className="text-center py-12">
              <AlertCircle className="mx-auto h-12 w-12 text-gray-400" />
              <h3 className="mt-2 text-sm font-medium text-gray-900">No clients</h3>
              <p className="mt-1 text-sm text-gray-500">Get started by creating a new client.</p>
              <div className="mt-6">
                <Button
                  onClick={() => setShowCreateModal(true)}
                  className="flex items-center mx-auto"
                >
                  <Plus className="h-5 w-5 mr-2" />
                  New Client
                </Button>
              </div>
            </div>
          ) : (
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>ID</TableHead>
                  <TableHead>Name</TableHead>
                  <TableHead>Created</TableHead>
                  <TableHead>Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {clients.map((client) => (
                  <TableRow key={client.id}>
                    <TableCell className="font-medium">{client.id}</TableCell>
                    <TableCell>{client.name}</TableCell>
                    <TableCell className="text-gray-500">
                      {new Date(client.created_at).toLocaleDateString()}
                    </TableCell>
                    <TableCell>
                      <div className="flex items-center space-x-2">
                        <Button
                          variant="ghost"
                          size="sm"
                          onClick={() => handleViewMappings(client)}
                        >
                          <Settings className="h-4 w-4 mr-1" />
                          Mappings
                        </Button>
                        <Button
                          variant="ghost"
                          size="sm"
                          onClick={() => handleDeleteClient(client.id, client.name)}
                          className="text-red-600 hover:text-red-900"
                        >
                          <Trash2 className="h-4 w-4 mr-1" />
                          Delete
                        </Button>
                      </div>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          )}
        </CardContent>
      </Card>

      {/* Create Client Modal */}
      <Dialog open={showCreateModal} onOpenChange={setShowCreateModal}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Create New Client</DialogTitle>
            <DialogDescription>
              Enter a name for the new client to get started.
            </DialogDescription>
          </DialogHeader>
          <form onSubmit={handleCreateClient}>
            <div className="space-y-4">
              <div>
                <label htmlFor="clientName" className="block text-sm font-medium text-gray-700 mb-2">
                  Client Name
                </label>
                <Input
                  id="clientName"
                  value={newClientName}
                  onChange={(e) => setNewClientName(e.target.value)}
                  placeholder="Enter client name"
                  required
                  disabled={creating}
                />
              </div>
            </div>
            <DialogFooter className="mt-6">
              <Button
                type="button"
                variant="outline"
                onClick={() => {
                  setShowCreateModal(false);
                  setNewClientName('');
                }}
                disabled={creating}
              >
                Cancel
              </Button>
              <Button
                type="submit"
                disabled={creating || !newClientName.trim()}
              >
                {creating ? 'Creating...' : 'Create Client'}
              </Button>
            </DialogFooter>
          </form>
        </DialogContent>
      </Dialog>

      {/* Mappings Modal */}
      <Dialog open={!!selectedClient} onOpenChange={() => setSelectedClient(null)}>
        <DialogContent className="max-w-6xl max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle>
              Mapping Rules for {selectedClient?.name}
            </DialogTitle>
            <DialogDescription>
              Manage mapping rules to transform data for this client.
            </DialogDescription>
          </DialogHeader>
          
          <div className="space-y-6">
            {/* Add New Mapping Button */}
            <div className="flex justify-between items-center">
              <h3 className="text-lg font-medium">Existing Mapping Rules</h3>
              <div className="space-x-2">
                <Button
                  variant="outline"
                  size="sm"
                  onClick={openExpressionHelp}
                >
                  <Info className="h-4 w-4 mr-1" />
                  Expression Help
                </Button>
                <Button
                  onClick={() => setShowMappingForm(true)}
                  size="sm"
                >
                  <Plus className="h-4 w-4 mr-1" />
                  Add Mapping Rule
                </Button>
              </div>
            </div>

            {/* New Mapping Form */}
            {showMappingForm && (
              <Card>
                <CardHeader>
                  <CardTitle className="flex items-center justify-between">
                    Create New Mapping Rule
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={() => setShowMappingForm(false)}
                    >
                      <X className="h-4 w-4" />
                    </Button>
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <form onSubmit={handleCreateMapping} className="space-y-4">
                    <div className="grid grid-cols-2 gap-4">
                      <div>
                        <label className="block text-sm font-medium mb-2">
                          Source Path <span className="text-red-500">*</span>
                        </label>
                        <Input
                          value={newMapping.source_path}
                          onChange={(e) => setNewMapping({...newMapping, source_path: e.target.value})}
                          placeholder="e.g., user.profile.name"
                          required
                        />
                        <p className="text-xs text-gray-500 mt-1">Dot notation path in source data</p>
                      </div>
                      <div>
                        <label className="block text-sm font-medium mb-2">
                          Destination Path <span className="text-red-500">*</span>
                        </label>
                        <Input
                          value={newMapping.destination_path}
                          onChange={(e) => setNewMapping({...newMapping, destination_path: e.target.value})}
                          placeholder="e.g., customer.fullName"
                          required
                        />
                        <p className="text-xs text-gray-500 mt-1">Dot notation path in output data</p>
                      </div>
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                      <div>
                        <label className="block text-sm font-medium mb-2">
                          Transform Type <span className="text-red-500">*</span>
                        </label>
                        <Select
                          value={newMapping.transform_type}
                          onValueChange={(value) => setNewMapping({...newMapping, transform_type: value})}
                        >
                          <SelectTrigger>
                            <SelectValue />
                          </SelectTrigger>
                          <SelectContent>
                            <SelectItem value="copy">Copy (direct mapping)</SelectItem>
                            <SelectItem value="toString">To String</SelectItem>
                            <SelectItem value="toBool">To Boolean</SelectItem>
                            <SelectItem value="toUpperCase">To Uppercase</SelectItem>
                            <SelectItem value="toLowerCase">To Lowercase</SelectItem>
                            <SelectItem value="capitalize">Capitalize</SelectItem>
                            <SelectItem value="formatDate">Format Date</SelectItem>
                            <SelectItem value="mapGender">Map Gender</SelectItem>
                            <SelectItem value="expression">Custom Expression</SelectItem>
                          </SelectContent>
                        </Select>
                      </div>
                      <div>
                        <label className="block text-sm font-medium mb-2">
                          Default Value
                        </label>
                        <Input
                          value={newMapping.default_value}
                          onChange={(e) => setNewMapping({...newMapping, default_value: e.target.value})}
                          placeholder="Value if source is missing"
                        />
                      </div>
                    </div>

                    {newMapping.transform_type === 'expression' && (
                      <div>
                        <label className="block text-sm font-medium mb-2">
                          Transform Logic <span className="text-red-500">*</span>
                          <Button
                            type="button"
                            variant="ghost"
                            size="sm"
                            onClick={openExpressionHelp}
                            className="ml-2"
                          >
                            <Code className="h-3 w-3 mr-1" />
                            Help
                          </Button>
                        </label>
                        <Textarea
                          value={newMapping.transform_logic}
                          onChange={(e) => setNewMapping({...newMapping, transform_logic: e.target.value})}
                          placeholder="e.g., toUpper(value) + ' - ' + getCurrentDate()"
                          rows={3}
                        />
                        <p className="text-xs text-gray-500 mt-1">JavaScript-like expression for transformation</p>
                      </div>
                    )}

                    <div className="flex items-center space-x-2">
                      <input
                        type="checkbox"
                        id="required"
                        checked={newMapping.required}
                        onChange={(e) => setNewMapping({...newMapping, required: e.target.checked})}
                        className="rounded"
                      />
                      <label htmlFor="required" className="text-sm">
                        Required field (transformation will fail if source is missing)
                      </label>
                    </div>

                    <div className="flex justify-end space-x-2">
                      <Button
                        type="button"
                        variant="outline"
                        onClick={() => setShowMappingForm(false)}
                        disabled={savingMapping}
                      >
                        Cancel
                      </Button>
                      <Button
                        type="submit"
                        disabled={savingMapping || !newMapping.source_path || !newMapping.destination_path}
                      >
                        {savingMapping ? (
                          <>
                            <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
                            Saving...
                          </>
                        ) : (
                          <>
                            <Save className="h-4 w-4 mr-1" />
                            Save Rule
                          </>
                        )}
                      </Button>
                    </div>
                  </form>
                </CardContent>
              </Card>
            )}
          
            {/* Existing Mappings Table */}
            {mappings.length === 0 ? (
              <div className="text-center py-8">
                <p className="text-gray-500">No mapping rules found for this client.</p>
                <p className="text-sm text-gray-400 mt-2">
                  Create mapping rules to enable data transformations.
                </p>
              </div>
            ) : (
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead>Source Path</TableHead>
                    <TableHead>Destination Path</TableHead>
                    <TableHead>Transform Type</TableHead>
                    <TableHead>Logic/Default</TableHead>
                    <TableHead>Required</TableHead>
                    <TableHead>Created</TableHead>
                    <TableHead>Actions</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {mappings.map((mapping) => (
                    <TableRow key={mapping.id}>
                      <TableCell>
                        <code className="text-xs bg-gray-100 px-2 py-1 rounded">
                          {mapping.source_path.join('.')}
                        </code>
                      </TableCell>
                      <TableCell>
                        <code className="text-xs bg-gray-100 px-2 py-1 rounded">
                          {mapping.destination_path.join('.')}
                        </code>
                      </TableCell>
                      <TableCell>
                        <Badge variant="secondary">
                          {mapping.transform_type}
                        </Badge>
                      </TableCell>
                      <TableCell className="max-w-48">
                        {mapping.transform_logic ? (
                          <code className="text-xs bg-blue-50 px-2 py-1 rounded block truncate">
                            {mapping.transform_logic}
                          </code>
                        ) : mapping.default_value ? (
                          <span className="text-xs text-gray-600">{mapping.default_value}</span>
                        ) : (
                          <span className="text-xs text-gray-400">None</span>
                        )}
                      </TableCell>
                      <TableCell>
                        {mapping.required ? (
                          <span className="text-green-600">Yes</span>
                        ) : (
                          <span className="text-gray-400">No</span>
                        )}
                      </TableCell>
                      <TableCell className="text-gray-500">
                        {new Date(mapping.created_at).toLocaleDateString()}
                      </TableCell>
                      <TableCell>
                        <Button
                          variant="ghost"
                          size="sm"
                          onClick={async () => {
                            if (confirm('Are you sure you want to delete this mapping rule?')) {
                              try {
                                await mappingAPI.delete(mapping.id);
                                toast.success('Mapping rule deleted');
                                loadMappings(selectedClient.id);
                              } catch (error) {
                                toast.error('Failed to delete mapping rule');
                              }
                            }
                          }}
                          className="text-red-600 hover:text-red-900"
                        >
                          <Trash2 className="h-4 w-4" />
                        </Button>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            )}
          </div>
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default Clients;
